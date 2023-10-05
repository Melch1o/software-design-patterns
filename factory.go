package main

import (
	"fmt"
)

// Interface for vehicle
type IVehicle interface {
	setType(string)
	getType() string
}

// Concrete vehicle
type Vehicle struct {
	vehicleType string
}

func (v *Vehicle) setType(s string) {
	v.vehicleType = s
}

func (v *Vehicle) getType() string {
	return v.vehicleType
}

// Concrete products
type car struct {
	Vehicle
}

func makeCar() IVehicle {
	return &car{
		Vehicle{
			vehicleType: "Car",
		},
	}
}

type plane struct {
	Vehicle
}

func makePlane() IVehicle {
	return &plane{
		Vehicle{
			vehicleType: "Plane",
		},
	}
}

type boat struct {
	Vehicle
}

func makeBoat() IVehicle {
	return &boat{
		Vehicle{
			vehicleType: "Boat",
		},
	}
}

func makeVehicle(vType string) IVehicle {
	if vType == "car" {
		return makeCar()
	}
	if vType == "plane" {
		return makePlane()
	}
	if vType == "boat" {
		return makeBoat()
	}
	return nil
}

func showInfo(v IVehicle) {
	fmt.Printf("Vehicle type: %s \n", v.getType())
}

func main() {
	someRandomBoatName := makeVehicle("boat")
	someCoolCar := makeVehicle("car")
	ifOnlyICouldAffordAPlane := makeVehicle("plane")

	showInfo(someRandomBoatName)
	showInfo(someCoolCar)
	showInfo(ifOnlyICouldAffordAPlane)
}
