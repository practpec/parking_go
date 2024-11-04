package models

import (
	"fmt"                        
	"fyne.io/fyne/v2"           
	"fyne.io/fyne/v2/canvas"     
	"fyne.io/fyne/v2/storage"    
	"math/rand"                  
	"sync"                      
	"time"                      
)


type Car struct {
	id          int
	parkingTime time.Duration  
	image       *canvas.Image  
	space       int            
	exitImage   *canvas.Image  
}

// Constructor de Carros
func NewCar(id int) *Car {
	carImages := []string{
        "./assets/yellow.png",
        "./assets/blue.png",
        "./assets/white.png",
        "./assets/convertible.png",
        "./assets/gray.png",
		"./assets/red.png",
		"./assets/green.png",
    }
    exitImages := []string{
        "./assets/yellow_exit.png",
        "./assets/blue_exit.png",
        "./assets/white_exit.png",
        "./assets/convertible_exit.png",
        "./assets/gray_exit.png",
		"./assets/red_exit.png",
		"./assets/green_exit.png",
    }

    randomIndex := rand.Intn(len(carImages))

    image := canvas.NewImageFromURI(storage.NewFileURI(carImages[randomIndex]))
    exitImage := canvas.NewImageFromURI(storage.NewFileURI(exitImages[randomIndex]))

	return &Car{
		id:          id,
		parkingTime: time.Duration(rand.Intn(5)+6) * time.Second,//duracion en el estacionamiento
		image:       image,
		space:       0,
		exitImage:   exitImage,
	}
}

var positions = [10]fyne.Position{
    {X: 415, Y: 15}, {X: 415, Y: 75}, {X: 415, Y: 135}, {X: 415, Y: 195}, {X: 415, Y: 255},
    {X: 205, Y: 15}, {X: 205, Y: 75}, {X: 205, Y: 135}, {X: 205, Y: 195}, {X: 205, Y: 255},
}

func (c *Car) Access(p *Parking, carsContainer *fyne.Container, statusText *canvas.Text) {
    p.SpacesChan() <- c.GetId()
    p.EntranceMutex().Lock()
    time.Sleep(300 * time.Millisecond)

    spacesArray := p.GetSpaces()
    message := fmt.Sprintf("%d	Dentro	%d", c.GetId(), len(p.SpacesChan()))
    statusText.Text = message
    statusText.Refresh()

    availablePositions := []int{}
    for i := 0; i < len(spacesArray); i++ {
        if !spacesArray[i] {
            availablePositions = append(availablePositions, i)
        }
    }

    if len(availablePositions) == 0 {
        message = fmt.Sprintf("No hay espacios disponibles %d", c.id)
        statusText.Text = message
        statusText.Refresh()
        p.EntranceMutex().Unlock()
        return
    }

    randomIndex := rand.Intn(len(availablePositions))
    spaceIndex := availablePositions[randomIndex]
    spacesArray[spaceIndex] = true
    c.space = spaceIndex 
    c.image.Move(positions[spaceIndex])

    p.SetSpaces(spacesArray)
    p.EntranceMutex().Unlock()
    carsContainer.Refresh()
}

func (c *Car) Goout(p *Parking, carsContainer *fyne.Container, statusText *canvas.Text) {
    p.EntranceMutex().Lock()
    <-p.SpacesChan()
    time.Sleep(300 * time.Millisecond) 

    spacesArray := p.GetSpaces()
    spacesArray[c.space] = false     
    p.SetSpaces(spacesArray) 

    message := fmt.Sprintf("%d	Fuera		%d", c.GetId(), len(p.SpacesChan()))
    statusText.Text = message
    statusText.Refresh()

    p.EntranceMutex().Unlock()

    for i := 0; i < 10; i++ {
        c.exitImage.Move(fyne.NewPos(c.exitImage.Position().X-30, c.exitImage.Position().Y))
        time.Sleep(time.Millisecond * 200)
    }

    carsContainer.Remove(c.exitImage)
    carsContainer.Refresh()
}

func (c *Car) Park(p *Parking, carsContainer *fyne.Container, wg *sync.WaitGroup, statusText *canvas.Text) {
    for i := 0; i < 7; i++ {
        c.image.Move(fyne.NewPos(c.image.Position().X+20, c.image.Position().Y))
        time.Sleep(time.Millisecond * 200)
    }

    c.Access(p, carsContainer, statusText)

    time.Sleep(c.parkingTime)

    carsContainer.Remove(c.image)
    c.exitImage.Resize(fyne.NewSize(65, 45))
    p.ExitCar(carsContainer, c.exitImage)
    c.Goout(p, carsContainer, statusText)

    wg.Done()
}

func (c *Car) GetId() int {
	return c.id
}

func (c *Car) GetCarImage() *canvas.Image {
	return c.image
}
