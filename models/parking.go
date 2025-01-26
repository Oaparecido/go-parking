package models

import (
	i "github.com/Oaparecido/go-parking/interfaces"
	"time"
)

type Parking struct {
	name   string
	spaces []Space
}

func (p Parking) GetName() string {
	return p.name
}

func (p *Parking) SetName(name string) {
	p.name = name
}

func (p *Parking) SetSpaces(qtd int) {
	p.spaces = make([]Space, qtd)

	for i := range p.spaces {
		p.spaces[i].free = true
		p.spaces[i].id = i + 1
	}
}

func (p Parking) GetSpaces() []Space {
	return p.spaces
}

func (p *Parking) ParkVehicle(idSpace int, vehicle i.Vehicle) {
	idSpace -= 1
	canPark := len(p.spaces) > 0 && p.spaces[idSpace].free == true

	if canPark {
		p.spaces[idSpace].id = idSpace + 1
		p.spaces[idSpace].free = false
		p.spaces[idSpace].vehicle = vehicle
		p.spaces[idSpace].vehicle.SetIncoming(time.Now())
	}
	// TODO: Adding if can't park vehicle in space.
}

func (p *Parking) RemoveCar(idSpace int, vehicle i.Vehicle) {
	idSpace -= 1
	canRemove := p.spaces[idSpace].free == false && p.spaces[idSpace].vehicle == vehicle

	if canRemove {
		p.spaces[idSpace].free = true
		p.spaces[idSpace].vehicle = nil
	}
}
