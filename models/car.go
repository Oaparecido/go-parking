package models

import (
	"time"
)

type Car struct {
	plate    string
	rate     float64
	incoming time.Time
}

func (c *Car) GetPlate() string {
	return c.plate
}

func (c *Car) GetRate() float64 {
	return c.rate
}

func (c *Car) GetIncoming() time.Time {
	return c.incoming
}

func (c *Car) SetIncoming(incoming time.Time) {
	c.incoming = incoming
}

func (c *Car) NewVehicle(plate string) {
	c.plate = plate
	c.rate = 3.5
}

func (c Car) ChargeRate(startTime, endTime time.Time) float64 {
	duration := endTime.Add(1 * time.Hour).Sub(startTime).Hours()
	return c.GetRate() * duration
}
