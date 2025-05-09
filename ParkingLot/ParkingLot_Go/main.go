package main

import (
	"fmt"
	"parkingLot/parkingLot"
	"parkingLot/vehicles"
)

func main() {
	parkingLot := parkingLot.GetNewParkingLot()

	parkingLot.AddLevel(10)

	truck1 := vehicles.GetNewTruck("1234")
	truck2 := vehicles.GetNewTruck("12345")
	b1 := vehicles.GetNewBike("231")
	c1 := vehicles.GetNewCar("321")

	fmt.Println(parkingLot.ParkVehicle(truck1))
	fmt.Println(parkingLot.ParkVehicle(truck2))
	fmt.Println(parkingLot.ParkVehicle(b1))
	fmt.Println(parkingLot.ParkVehicle(c1))

	fmt.Println(parkingLot.CheckAvailability())

	fmt.Println(parkingLot.UnParkVehicle(truck2))
	fmt.Println(parkingLot.UnParkVehicle(truck1))

	fmt.Println(parkingLot.CheckAvailability())
}
