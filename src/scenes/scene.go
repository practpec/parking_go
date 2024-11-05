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


	parking := canvas.NewText("No.Auto     Estado     Esp.Ocupado", color.White)
	parking.TextSize = 10
	parking.Move(fyne.NewPos(10, 10))
	parking.TextStyle = fyne.TextStyle{Bold: true}

    s.statusText.Move(fyne.NewPos(20, 20))
    carsContainer.Move(fyne.NewPos(0, 0))
    s.window.SetContent(container.NewWithoutLayout(font, rectangle, road, line, line2, gate,parking, textGroup, s.statusText, carsContainer))
}


func (s *Scene) Run() {
    p := models.NewParking(make(chan int, 10), &sync.Mutex{})
    randomGen := models.NewRandom()
    var wg sync.WaitGroup
    creationLimiter := make(chan struct{}, 1)

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go s.createCar(i, p, &wg, creationLimiter, randomGen)
    }

    wg.Wait()
    fmt.Println("Acabo")
}


func (s *Scene) createCar(id int, p *models.Parking, wg *sync.WaitGroup, creationLimiter chan struct{}, randomGen *models.Random) {
    creationLimiter <- struct{}{}
    defer func() { <-creationLimiter }()

    car := models.NewCar(id)
    carImage := car.GetCarImage()
    carImage.Resize(fyne.NewSize(65, 45))
    carImage.Move(fyne.NewPos(-20, 310))

    carsContainer.Add(carImage)
    carsContainer.Refresh()


    go s.parkCar(car, p, carsContainer, wg)

    var randNumber = randomGen.Generate(float64(1))
    time.Sleep(time.Second * time.Duration(randNumber))
}


func (s *Scene) parkCar(car *models.Car, p *models.Parking, carsContainer *fyne.Container, wg *sync.WaitGroup) {
    car.Park(p, carsContainer, wg, s.statusText)
}



