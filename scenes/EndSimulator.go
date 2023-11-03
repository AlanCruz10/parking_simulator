package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type EndSimulator struct {
	window    fyne.Window
	app       fyne.App
	container *fyne.Container
}

func NewEndSimulator(w fyne.Window, app fyne.App) *EndSimulator {
	return &EndSimulator{
		window:    w,
		app:       app,
		container: container.NewWithoutLayout(),
	}
}

func (e *EndSimulator) EndSimulatorScene() {

	backgroundGameOver := canvas.NewImageFromURI(storage.NewFileURI("./assets/backgrounds/simulator_finished.png"))
	backgroundGameOver.Resize(fyne.NewSize(1240, 720))
	backgroundGameOver.Move(fyne.NewPos(0, 75))
	e.container.Add(backgroundGameOver)

	//menu := widget.NewButton("Volver al menú", func() {
	//	e.BackMenu()
	//})
	//menu.Resize(fyne.NewSize(200, 50))
	//menu.Move(fyne.NewPos(400, 450))
	//e.container.Add(menu)
	//
	//restart := widget.NewButton("Volver a simular", func() {
	//	e.Restart()
	//})
	//restart.Resize(fyne.NewSize(200, 50))
	//restart.Move(fyne.NewPos(650, 450))
	//e.container.Add(restart)

	exit := widget.NewButton("Salir", e.Exit)
	exit.Resize(fyne.NewSize(200, 50))
	exit.Move(fyne.NewPos(525, 550))
	e.container.Add(exit)

	e.window.SetContent(e.container)

}

//func (e *EndSimulator) BackMenu() {
//	scene := NewMenuScene(e.window, e.app)
//	scene.Menu()
//}
//
//func (e *EndSimulator) Restart() {
//	simulator := NewSimulator(e.window, e.app)
//	simulator.StartSimulator()
//}

func (e *EndSimulator) Exit() {
	e.window.Close()
}
