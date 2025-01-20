package interfaces

import "time"

type Vehicle interface {
	GetPlate() string
	GetRate() float64
	GetIncoming() time.Time
	SetIncoming(incoming time.Time)
	NewVehicle(plate string)
	ChargeRate(timeStart, timeEnd time.Time) float64
}
