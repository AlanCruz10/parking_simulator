package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type MainWindow struct {
	App    fyne.App
	Window fyne.Window
}

func NewMainWindow(title string, size fyne.Size) *MainWindow {
	mainApp := app.New()
	mainWindow := mainApp.NewWindow(title)
	mainWindow.CenterOnScreen()
	mainWindow.SetFixedSize(true)
	mainWindow.Resize(size)
	return &MainWindow{
		App:    mainApp,
		Window: mainWindow,
	}
}

func (w *MainWindow) StarSimulator() {
	menuScene := NewMenuScene(w.Window, w.App)
	menuScene.Menu()
	w.Window.ShowAndRun()
}
