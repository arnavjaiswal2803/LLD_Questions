package level

import (
	"parkingLot/parkingSpot"
	"parkingLot/vehicles"
	"sync"
)

func GetNewLevel(floor, spotsCount int) *Level {
	parkingSpots := make([]*parkingSpot.ParkingSpot, spotsCount)

	carSpots := int(float32(spotsCount) * 0.6)
	for i := 0; i < carSpots; i++ {
		parkingSpots[i] = parkingSpot.GetNewParkingSpot(i, vehicles.Car)
	}

	bikeSpots := int(float32(spotsCount) * 0.3)
	for i := carSpots; i < carSpots+bikeSpots; i++ {
		parkingSpots[i] = parkingSpot.GetNewParkingSpot(i, vehicles.Bike)
	}

	for i := carSpots + bikeSpots; i < spotsCount; i++ {
		parkingSpots[i] = parkingSpot.GetNewParkingSpot(i, vehicles.Truck)
	}

	return &Level{
		floor:        floor,
		parkingSpots: parkingSpots,
	}
}

type Level struct {
	floor        int
	parkingSpots []*parkingSpot.ParkingSpot
	mu           sync.Mutex
}

func (l *Level) GetId() int {
	return l.floor
}

func (l *Level) ParkVehicle(v vehicles.Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if spot.GetType() == v.GetType() && spot.IsAvailable(&l.mu) {
			spot.ParkVehicle(v, &l.mu)
			return true
		}
	}
	return false
}

func (l *Level) UnParkVehicle(v vehicles.Vehicle) bool {
	for _, spot := range l.parkingSpots {
		if spot.GetVehicle() != nil && spot.GetVehicle().GetLicense() == v.GetLicense() {
			spot.UnParkVehicle(&l.mu)
			return true
		}
	}
	return false
}

func (l *Level) CheckAvailability() int {
	count := 0
	for _, spot := range l.parkingSpots {
		if spot.IsAvailable(&l.mu) {
			count++
		}
	}
	return count
}
