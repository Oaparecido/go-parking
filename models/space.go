package models

import (
	i "github.com/Oaparecido/go-parking/interfaces"
)

type Space struct {
	id      int
	vehicle i.Vehicle
	free    bool
}

func (s Space) GetSpace() (int, i.Vehicle, bool) {
	return s.id, s.vehicle, s.free
}
