package views

import (
	"parking_go/src/scenes"
	"fyne.io/fyne/v2"      
	"fyne.io/fyne/v2/app"   
)


type View struct{}


func NewView() *View {
	return &View{}
}

func (v *View) Run() {
	myApp := app.New()
	window := myApp.NewWindow("Estacionamiento Concurrente")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(500, 420))
	scene := scenes.NewScene(window)
	scene.Show()
	go scene.Run()
	window.ShowAndRun()
}
