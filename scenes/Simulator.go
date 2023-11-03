package scenes

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"math/rand"
	"parking_simulator/models"
	"time"
)

type Simulator struct {
	Window    fyne.Window
	App       fyne.App
	Container *fyne.Container
}

func NewSimulator(window fyne.Window, app fyne.App) *Simulator {
	return &Simulator{
		Window:    window,
		App:       app,
		Container: container.NewWithoutLayout(),
	}
}

func (s *Simulator) StartSimulator() {
	parking := models.NewParking(s.Window, s.App, s.Container, 1240, 720)
	parking.GetImageParking().Resize(parking.GetSize())
	parking.GetImageParking().Move(parking.GetPosition())
	s.Container.Add(parking.GetImageParking())

	parking.CounterLabel = widget.NewLabel(fmt.Sprintf("Vehicles: %d ", parking.ArrivedCount))
	parking.CounterLabel.TextStyle = fyne.TextStyle{Bold: true, TabWidth: 0}
	parking.CounterLabel.Resize(fyne.NewSize(200, 400))
	parking.CounterLabel.Move(fyne.NewPos(200, 100))
	s.Container.Add(parking.CounterLabel)

	s.Window.SetContent(s.Container)

	go s.Star(parking)
	go s.Count(parking)

}

func (s *Simulator) Star(parking *models.Parking) {
	for i := 0; i < 100; i++ {
		go parking.CreateVehicles()
		time.Sleep(time.Duration(poissonTimeArrived()) * time.Second)
	}
}

func (s *Simulator) Count(parking *models.Parking) {
	for {
		select {
		case count := <-parking.CounterUpdate:
			parking.CounterLabel.SetText(fmt.Sprintf("Vehicles: %d", count))
			if count == 100 {
				endSimulator := NewEndSimulator(s.Window, s.App)
				endSimulator.EndSimulatorScene()
			}
		}
	}
}

func poissonTimeArrived() int {
	u := rand.Float64()
	interval := -math.Log(1.0-u) / 0.5
	return int(interval)
}
