package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MenuScene struct {
	window fyne.Window
	app    fyne.App
}

func NewMenuScene(w fyne.Window, a fyne.App) *MenuScene {
	return &MenuScene{
		window: w,
		app:    a,
	}
}

func (m *MenuScene) Menu() {

	backgroundTitle := canvas.NewImageFromURI(storage.NewFileURI("./assets/backgrounds/parking_simulator.png"))
	backgroundTitle.Resize(fyne.NewSize(1240, 720))
	backgroundTitle.Move(fyne.NewPos(0, 0))

	start := widget.NewButton("START", m.Start)
	start.Resize(fyne.NewSize(200, 50))
	start.Move(fyne.NewPos(150, 300))

	exit := widget.NewButton("EXIT", m.Exit)
	exit.Resize(fyne.NewSize(200, 50))
	exit.Move(fyne.NewPos(900, 300))

	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/backgrounds/car_menu.jpg"))
	backgroundImage.Resize(fyne.NewSize(1240, 720))
	backgroundImage.Move(fyne.NewPos(0, 0))

	m.window.SetContent(container.NewWithoutLayout(backgroundImage, backgroundTitle, start, exit))

}

func (m *MenuScene) Start() {
	simulator := NewSimulator(m.window, m.app)
	simulator.StartSimulator()
}

func (m *MenuScene) Exit() {
	m.window.Close()
}
