package main

import (
	"fmt"
	i "github.com/Oaparecido/go-parking/interfaces"
	m "github.com/Oaparecido/go-parking/models"
	"os"
	"time"
)

var p m.Parking
var v i.Vehicle

var parks []m.Parking
var vehicles []i.Vehicle

func main() {
	fmt.Println("Starting service...")
	time.Sleep(2 * time.Second)
	fmt.Println()

	for {
		fmt.Println("**** Parking station ****")
		fmt.Println("*                       *")
		fmt.Println("* A -> Register park    *")
		fmt.Println("* B -> Register vehicle *")
		fmt.Println("* C -> Park vehicle     *")
		fmt.Println("* D -> See p & v        *")
		fmt.Println("* E -> Exit             *")
		fmt.Println("*                       *")
		fmt.Println("*************************")
		fmt.Println()

		var command string
		fmt.Scan(&command)

		switch command {
		case "A":
			registerPark(p)
		case "B":
			registerVehicle()
		case "C":
			parkVehicle()
		case "D":
			seeParksAndVehicles()

		case "E":
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}

func registerPark(p m.Parking) {
	fmt.Println("* -> Name:")
	var parkName string
	fmt.Scan(&parkName)

	fmt.Println("* -> Quantity spaces:")
	var spacesQtd int
	fmt.Scan(&spacesQtd)

	p.SetName(parkName)
	p.SetSpaces(spacesQtd)

	parks = append(parks, p)

	fmt.Println()
	fmt.Println(parks)
	fmt.Println()
}

func registerVehicle() {
	var command string

	fmt.Println("**** Select vehicle type ****")
	fmt.Println("*                           *")
	fmt.Println("* A -> Car                  *")
	fmt.Println("* B -> Scooter              *")
	fmt.Println("*                           *")
	fmt.Println("*****************************")
	fmt.Println()

	fmt.Scan(&command)

	switch command {
	case "A":
		v = &m.Car{}
	}

	fmt.Println("* -> Plate:")
	var vPlate string
	fmt.Scan(&vPlate)

	v.NewVehicle(vPlate)
	vehicles = append(vehicles, v)
	fmt.Println()
	fmt.Println(vehicles)
	fmt.Println()
}

func parkVehicle() {
	fmt.Println("**** Select vehicle ****")

	for i, v := range vehicles {
		fmt.Println("* ", i, "->", v.GetPlate())
	}

	fmt.Println("************************")

	var index int

	fmt.Scan(&index)

	vChoosen := vehicles[index]

	fmt.Println("* -> Choosen vehicle:", vChoosen.GetPlate())
	fmt.Println()

	fmt.Println("**** Select park ****")

	for i, p := range parks {
		fmt.Println("* ", i, "->", p.GetName())
	}

	fmt.Println("************************")

	fmt.Scan(&index)

	pChoosen := parks[index]

	fmt.Println("* -> Choosen park:", pChoosen.GetName())
	fmt.Println()

	pChoosen.ParkVehicle(1, vChoosen)
	// TODO: Add validation if not park vehicle, show spaces free
}

func seeParksAndVehicles() {
	fmt.Println("* -> Parks", parks)
	fmt.Println("* -> Vehicles", vehicles)
	fmt.Println()
}
