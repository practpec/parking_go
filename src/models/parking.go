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
    observers   []Observer 
}

func NewParking(spaces chan int, entrance *sync.Mutex) *Parking {
    return &Parking{
        spaces:      spaces,
        entrance:    entrance,
        spacesArray: [10]bool{},
        observers:   []Observer{},
    }
}


func (p *Parking) RegisterObserver(observer Observer) {
    p.observers = append(p.observers, observer)
}

func (p *Parking) NotifyObservers(spaceAvailable bool) {
    for _, observer := range p.observers {
        observer.Update(spaceAvailable)
    }
}


func (p *Parking) ExitCar(carsContainer *fyne.Container, carImage *canvas.Image) {
    carImage.Move(fyne.NewPos(205, 350))
    carsContainer.Add(carImage)
    carsContainer.Refresh()

   
    p.NotifyObservers(true)
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


