package scenes

import (
	"parking_go/src/models"
	"fmt"                   
	"fyne.io/fyne/v2"       
	"fyne.io/fyne/v2/canvas" 
	"fyne.io/fyne/v2/container"
	"image/color"           
	"sync"                  
	"time"                  
)


type Scene struct {
	window     fyne.Window 
	statusText *canvas.Text 
}

func NewScene(window fyne.Window) *Scene {
	statusText := canvas.NewText("", color.White) 
	return &Scene{
		window:     window,
		statusText: statusText,
	}
}

var carsContainer = container.NewWithoutLayout()

// Elementos Visuales
func (s *Scene) Show() {

	font := canvas.NewRectangle(color.RGBA{R: 0, G: 100, B: 0, A: 255})
	font.Resize(fyne.NewSize(500, 420))
	font.Move(fyne.NewPos(0, 0))       

	rectangle := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 255})
	rectangle.StrokeWidth = 5 
	rectangle.StrokeColor = color.White
	rectangle.Resize(fyne.NewSize(300, 420))
	rectangle.Move(fyne.NewPos(200, 0))

	road := canvas.NewRectangle(color.RGBA{R: 128, G: 128, B: 128, A: 255})
	road.Resize(fyne.NewSize(195, 120))
	road.Move(fyne.NewPos(0, 300))

	line := canvas.NewRectangle(color.White)
	line.Resize(fyne.NewSize(350, 5))
	line.Move(fyne.NewPos(0, 360))

	line2 := canvas.NewRectangle(color.White)
	line2.Resize(fyne.NewSize(5, 365))
	line2.Move(fyne.NewPos(350, 0))

	gate := canvas.NewRectangle(color.White)
	gate.Resize(fyne.NewSize(10, 120))
	gate.Move(fyne.NewPos(195, 300))

	textGroup := container.NewWithoutLayout()
	text := "PARKING ZONE"
	for i, letter := range text {
		letterText := canvas.NewText(string(letter), color.White)
		letterText.TextSize = 20
		letterText.Move(fyne.NewPos(355, 20+float32(i*20)))
		textGroup.Add(letterText)
	}

	s.statusText.Move(fyne.NewPos(10, 20)) 
	parking := canvas.NewText("No.Auto     Estado     Esp.Ocupado", color.White)
	parking.TextSize = 10
	parking.Move(fyne.NewPos(10, 10))
	parking.TextStyle = fyne.TextStyle{Bold: true}
	carsContainer.Add(font)
	carsContainer.Add(rectangle)
	carsContainer.Add(road)
	carsContainer.Add(line)
	carsContainer.Add(line2)
	carsContainer.Add(gate)
	carsContainer.Add(textGroup)
	carsContainer.Add(parking)
	carsContainer.Add(s.statusText)
	s.window.SetContent(carsContainer)
}

//Elementos para correr el programa
func (s *Scene) Run() {
	p := models.NewParking(make(chan int, 10), &sync.Mutex{})
	randomGen := models.NewRandom()
	var wg sync.WaitGroup
	creationLimiter := make(chan struct{}, 1) // Canal para limitar la creación de coches

	for i := 0; i < 100; i++ {
		wg.Add(1)

		// Crear carro en un goroutine separado
		go func(id int) {
			creationLimiter <- struct{}{}
			defer func() { <-creationLimiter }()

			car := models.NewCar(id)
			carImage := car.GetCarImage()
			carImage.Resize(fyne.NewSize(65, 45))
			carImage.Move(fyne.NewPos(-20, 310))

			carsContainer.Add(carImage)
			carsContainer.Refresh()

			// Aparc del carro en otro goroutine para no bloquear la creación
			go func() {
				car.Park(p, carsContainer, &wg, s.statusText)
			}()

		}(i)

		var randNumber = randomGen.Generate(float64(1))
		time.Sleep(time.Second * time.Duration(randNumber))
	}

	wg.Wait()
	fmt.Println("Acabo")
}


