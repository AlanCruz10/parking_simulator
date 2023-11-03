package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"sync"
	"time"
)

const (
	capacityPark = 20
)

type Parking struct {
	Window            fyne.Window
	App               fyne.App
	Container         *fyne.Container
	ImageParking      *canvas.Image
	Size              fyne.Size
	Position          fyne.Position
	ArrivedCount      int
	Capacity          int
	Occupied          int
	SpacesStatusPark  [capacityPark]bool
	SpacePositionPark []fyne.Position
	Gate              sync.Mutex
	CounterLabel      *widget.Label
	CounterUpdate     chan int
}

func NewParking(window fyne.Window, app fyne.App, container *fyne.Container, sizeWith float32, sizeHeight float32) *Parking {
	parking := &Parking{
		Window:       window,
		App:          app,
		Container:    container,
		ImageParking: canvas.NewImageFromFile("assets/parking/parking_pro_two.png"),
		Size:         fyne.NewSize(sizeWith, sizeHeight),
		Position:     fyne.NewPos(0, 0),
		ArrivedCount: 0,
		Occupied:     0,
		Capacity:     capacityPark,
		SpacesStatusPark: [capacityPark]bool{
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
			false,
		},
		SpacePositionPark: []fyne.Position{
			fyne.NewPos(635, 100),
			fyne.NewPos(750, 100),
			fyne.NewPos(850, 100),
			fyne.NewPos(955, 100),
			fyne.NewPos(1070, 100),
			fyne.NewPos(635, 220),
			fyne.NewPos(750, 220),
			fyne.NewPos(850, 220),
			fyne.NewPos(955, 220),
			fyne.NewPos(1070, 220),
			fyne.NewPos(635, 410),
			fyne.NewPos(750, 410),
			fyne.NewPos(850, 410),
			fyne.NewPos(955, 410),
			fyne.NewPos(1070, 410),
			fyne.NewPos(635, 530),
			fyne.NewPos(750, 530),
			fyne.NewPos(850, 530),
			fyne.NewPos(955, 530),
			fyne.NewPos(1070, 530),
		},
		CounterUpdate: make(chan int),
	}
	return parking
}

func (p *Parking) GetImageParking() *canvas.Image {
	return p.ImageParking
}

func (p *Parking) GetPosition() fyne.Position {
	return p.Position
}

func (p *Parking) GetSize() fyne.Size {
	return p.Size
}

func (p *Parking) CreateVehicles() {
	vehicle := NewVehicle()
	vehicle.GetImageArrived().Resize(vehicle.GetSize())
	vehicle.GetImageArrived().Move(vehicle.GetPosition())
	p.Container.Add(vehicle.GetImageArrived())
	p.Container.Refresh()
	p.Enter(vehicle)
	p.ArrivedCount++
	p.CounterUpdate <- p.ArrivedCount
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	p.Exit(vehicle)
}

func (p *Parking) Enter(vehicle *Vehicle) {
	p.Gate.Lock()
	defer p.Gate.Unlock()
	p.ValidationSpacePark()
	for i := 0; i < p.Capacity; i++ {
		if !p.SpacesStatusPark[i] {
			p.SpacesStatusPark[i] = true
			p.Occupied++
			if i < 5 {
				x := float32(0)
				for j := vehicle.GetPosition().X; j < vehicle.GetPosition().X+500; j += 2 {
					vehicle.GetImageArrived().Move(fyne.NewPos(j, vehicle.GetPosition().Y))
					time.Sleep(time.Millisecond)
					x = j
				}
				vehicle.SetPosition(fyne.NewPos(x, vehicle.GetPosition().Y))
				p.Container.Remove(vehicle.GetImageArrived())
				vehicle.GetImageUp().Resize(vehicle.GetSize())
				vehicle.GetImageUp().Move(fyne.NewPos(vehicle.GetPosition().X, vehicle.GetPosition().Y))
				p.Container.Add(vehicle.GetImageUp())
				p.Container.Refresh()
				y := float32(0)
				for j := vehicle.GetPosition().Y; j >= p.SpacePositionPark[i].Y-75; j -= 2 {
					vehicle.GetImageUp().Move(fyne.NewPos(vehicle.GetPosition().X, j))
					time.Sleep(time.Millisecond)
					y = j
				}
				vehicle.SetPosition(fyne.NewPos(vehicle.GetPosition().X, y))
				p.Container.Remove(vehicle.GetImageUp())
				vehicle.GetImageArrived().Resize(vehicle.GetSize())
				vehicle.GetImageArrived().Move(fyne.NewPos(vehicle.GetPosition().X, vehicle.GetPosition().Y))
				p.Container.Add(vehicle.GetImageArrived())
				p.Container.Refresh()
				for j := vehicle.GetPosition().X; j < p.SpacePositionPark[i].X; j += 2 {
					vehicle.GetImageArrived().Move(fyne.NewPos(j, vehicle.GetPosition().Y))
					time.Sleep(time.Millisecond)
				}
				vehicle.SetPosition(fyne.NewPos(p.SpacePositionPark[i].X, p.SpacePositionPark[i].Y))
				p.Container.Remove(vehicle.GetImageArrived())
				vehicle.GetImageDown().Resize(vehicle.GetSize())
				vehicle.GetImageDown().Move(vehicle.GetPosition())
				p.Container.Add(vehicle.GetImageDown())
				p.Container.Refresh()
				vehicle.SetIdParkingSite(i)
				break
			} else if i >= 5 && i < 10 {
				for j := vehicle.GetPosition().X; j < p.SpacePositionPark[i].X; j += 2 {
					vehicle.GetImageArrived().Move(fyne.NewPos(j, vehicle.GetPosition().Y))
					time.Sleep(time.Millisecond)
				}
				p.Container.Remove(vehicle.GetImageArrived())
				vehicle.GetImageUp().Resize(vehicle.GetSize())
				vehicle.GetImageUp().Move(fyne.NewPos(p.SpacePositionPark[i].X, p.SpacePositionPark[i].Y))
				p.Container.Add(vehicle.GetImageUp())
				p.Container.Refresh()
				vehicle.SetIdParkingSite(i)
				break
			} else if i >= 10 && i < 15 {
				for j := vehicle.GetPosition().X; j < p.SpacePositionPark[i].X; j += 2 {
					vehicle.GetImageArrived().Move(fyne.NewPos(j, vehicle.GetPosition().Y))
					time.Sleep(time.Millisecond)
				}
				p.Container.Remove(vehicle.GetImageArrived())
				vehicle.GetImageDown().Resize(vehicle.GetSize())
				vehicle.GetImageDown().Move(fyne.NewPos(p.SpacePositionPark[i].X, p.SpacePositionPark[i].Y))
				p.Container.Add(vehicle.GetImageDown())
				p.Container.Refresh()
				vehicle.SetIdParkingSite(i)
				break
			}
			x := float32(0)
			for j := vehicle.GetPosition().X; j < vehicle.GetPosition().X+500; j += 2 {
				vehicle.GetImageArrived().Move(fyne.NewPos(j, vehicle.GetPosition().Y))
				time.Sleep(time.Millisecond)
				x = j
			}
			vehicle.SetPosition(fyne.NewPos(x, vehicle.GetPosition().Y))
			p.Container.Remove(vehicle.GetImageArrived())
			vehicle.GetImageDown().Resize(vehicle.GetSize())
			vehicle.GetImageDown().Move(fyne.NewPos(vehicle.GetPosition().X, vehicle.GetPosition().Y))
			p.Container.Add(vehicle.GetImageDown())
			p.Container.Refresh()
			y := float32(0)
			for j := vehicle.GetPosition().Y; j <= p.SpacePositionPark[i].Y+75; j += 2 {
				vehicle.GetImageDown().Move(fyne.NewPos(vehicle.GetPosition().X, j))
				time.Sleep(time.Millisecond)
				y = j
			}
			vehicle.SetPosition(fyne.NewPos(vehicle.GetPosition().X, y))
			p.Container.Remove(vehicle.GetImageDown())
			vehicle.GetImageArrived().Resize(vehicle.GetSize())
			vehicle.GetImageArrived().Move(fyne.NewPos(vehicle.GetPosition().X, vehicle.GetPosition().Y))
			p.Container.Add(vehicle.GetImageArrived())
			p.Container.Refresh()
			for j := vehicle.GetPosition().X; j < p.SpacePositionPark[i].X; j += 2 {
				vehicle.GetImageArrived().Move(fyne.NewPos(j, vehicle.GetPosition().Y))
				time.Sleep(time.Millisecond)
			}
			vehicle.SetPosition(fyne.NewPos(p.SpacePositionPark[i].X, p.SpacePositionPark[i].Y))
			p.Container.Remove(vehicle.GetImageArrived())
			vehicle.GetImageUp().Resize(vehicle.GetSize())
			vehicle.GetImageUp().Move(vehicle.GetPosition())
			p.Container.Add(vehicle.GetImageUp())
			p.Container.Refresh()
			vehicle.SetIdParkingSite(i)
			break
		}
	}
}

func (p *Parking) Exit(vehicle *Vehicle) {
	p.Gate.Lock()
	defer p.Gate.Unlock()
	for i := p.Capacity - 1; i >= 0; i-- {
		if i == vehicle.GetIdParkSite() {
			p.SpacesStatusPark[i] = false
			p.Occupied--
			if i < 5 {
				for j := p.SpacePositionPark[i].Y; j > p.SpacePositionPark[i].Y-100; j -= 2 {
					vehicle.GetImageDown().Move(fyne.NewPos(p.SpacePositionPark[i].X, j))
					time.Sleep(time.Millisecond)
				}
				p.Container.Remove(vehicle.GetImageDown())
				p.Container.Refresh()
				break
			} else if i >= 5 && i < 10 {
				for j := p.SpacePositionPark[i].Y; j < p.SpacePositionPark[i].Y+100; j += 2 {
					vehicle.GetImageUp().Move(fyne.NewPos(p.SpacePositionPark[i].X, j))
					time.Sleep(time.Millisecond)
				}
				p.Container.Remove(vehicle.GetImageUp())
				p.Container.Refresh()
				break
			} else if i >= 10 && i < 15 {
				for j := p.SpacePositionPark[i].Y; j > p.SpacePositionPark[i].Y-100; j -= 2 {
					vehicle.GetImageDown().Move(fyne.NewPos(p.SpacePositionPark[i].X, j))
					time.Sleep(time.Millisecond)
				}
				p.Container.Remove(vehicle.GetImageDown())
				p.Container.Refresh()
				break
			}
			for j := p.SpacePositionPark[i].Y; j < p.SpacePositionPark[i].Y+100; j += 2 {
				vehicle.GetImageUp().Move(fyne.NewPos(p.SpacePositionPark[i].X, j))
				time.Sleep(time.Millisecond)
			}
			p.Container.Remove(vehicle.GetImageUp())
			p.Container.Refresh()
			break
		}
	}
}

func (p *Parking) ValidationSpacePark() {
	time.Sleep(time.Second)
	for {
		if p.Occupied < p.Capacity {
			break
		}
		return
	}
}
