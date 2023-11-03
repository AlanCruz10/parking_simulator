package main

import (
	"fyne.io/fyne/v2"
	"parking_simulator/scenes"
)

func main() {
	mainWindow := scenes.NewMainWindow("PARKING SIMULATOR", fyne.NewSize(1240, 720))
	mainWindow.StarSimulator()
}
