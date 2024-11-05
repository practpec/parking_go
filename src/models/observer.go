package models

// Definición de la interfaz Observer
type Observer interface {
    Update(spaceAvailable bool)
}
