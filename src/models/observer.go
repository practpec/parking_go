package models

type Observer interface {
    Update(spaceAvailable bool)
}
