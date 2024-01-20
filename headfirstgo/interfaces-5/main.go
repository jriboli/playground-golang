package main

import "fmt"

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding Up")
}

func (t Truck) Brake() {
	fmt.Println("Braking")
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("right")
	vehicle.Steer("left")
	vehicle.Brake()
	if truck, ok := vehicle.(Truck); ok {
		truck.LoadCargo("pants")
	}
}

func main() {
	TryVehicle(Truck("Ford F180"))
}