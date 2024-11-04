package models


import (
	"fyne.io/fyne/v2"        
	"fyne.io/fyne/v2/canvas" 
	"sync"                   
)

type Parking struct {
	spaces      chan int    
	entrance    *sync.Mutex 
	spacesArray [10]bool    
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
	return &Parking{
		spaces:      spaces,             
		entrance:    entrance,           
		spacesArray: [10]bool{},         
	}
}

func (p *Parking) EntranceMutex() *sync.Mutex {
	return p.entrance
}

func (p *Parking) SpacesChan() chan int {
	return p.spaces
}

func (p *Parking) GetSpaces() [10]bool {
	return p.spacesArray
}

func (p *Parking) SetSpaces(spacesArray [10]bool) {
	p.spacesArray = spacesArray
}

func (p *Parking) ExitCar(carsContainer *fyne.Container, carImage *canvas.Image) {
	carImage.Move(fyne.NewPos(205, 350))
	carsContainer.Add(carImage)
	carsContainer.Refresh()
}
