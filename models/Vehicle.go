package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"time"
)

type Vehicle struct {
	ID            int
	IdParkingSite int
	Status        bool
	Position      fyne.Position
	Time          time.Duration
	Size          fyne.Size
	ImageArrived  *canvas.Image
	ImageLeft     *canvas.Image
	ImageDown     *canvas.Image
	ImageUp       *canvas.Image
	PositionX     int
	PositionY     int
}

func NewVehicle() *Vehicle {
	return &Vehicle{
		ID:            0,
		IdParkingSite: -1,
		Status:        false,
		Position:      fyne.NewPos(0, 317),
		Time:          time.Duration(0) * time.Second,
		Size:          fyne.NewSize(100, 100),
		ImageArrived:  canvas.NewImageFromFile("assets/cars/car_arrived.png"),
		ImageLeft:     canvas.NewImageFromFile("assets/cars/car_left.png"),
		ImageDown:     canvas.NewImageFromFile("assets/cars/car_down.png"),
		ImageUp:       canvas.NewImageFromFile("assets/cars/car_up.png"),
		PositionY:     0,
		PositionX:     0,
	}
}

func (v *Vehicle) GetID() int {
	return v.ID
}

func (v *Vehicle) GetIdParkSite() int {
	return v.IdParkingSite
}

func (v *Vehicle) GetStatus() bool {
	return v.Status
}

func (v *Vehicle) GetPosition() fyne.Position {
	return v.Position
}

func (v *Vehicle) GetTime() time.Duration {
	return v.Time
}

func (v *Vehicle) GetSize() fyne.Size {
	return v.Size
}

func (v *Vehicle) GetImageArrived() *canvas.Image {
	return v.ImageArrived
}

func (v *Vehicle) GetImageLeft() *canvas.Image {
	return v.ImageLeft
}

func (v *Vehicle) GetImageDown() *canvas.Image {
	return v.ImageDown
}

func (v *Vehicle) GetImageUp() *canvas.Image {
	return v.ImageUp
}

func (v *Vehicle) GetPositionX() int {
	return v.PositionX
}

func (v *Vehicle) GetPositionY() int {
	return v.PositionY
}

func (v *Vehicle) SetID(id int) {
	v.ID = id
}

func (v *Vehicle) SetIdParkingSite(idParkingSite int) {
	v.IdParkingSite = idParkingSite
}

func (v *Vehicle) SetStatus(status bool) {
	v.Status = status
}

func (v *Vehicle) SetPosition(position fyne.Position) {
	v.Position = position
}

func (v *Vehicle) SetTime(time time.Duration) {
	v.Time = time
}

func (v *Vehicle) SetImageArrived(imageArrived *canvas.Image) {
	v.ImageArrived = imageArrived
}

func (v *Vehicle) SetImageLeft(imageLeft *canvas.Image) {
	v.ImageLeft = imageLeft
}

func (v *Vehicle) SetImageDown(imageDown *canvas.Image) {
	v.ImageDown = imageDown
}

func (v *Vehicle) SetImageUp(imageUp *canvas.Image) {
	v.ImageUp = imageUp
}

func (v *Vehicle) SetPositionX(positionX int) {
	v.PositionX = positionX
}

func (v *Vehicle) SetPositionY(positionY int) {
	v.PositionY = positionY
}
