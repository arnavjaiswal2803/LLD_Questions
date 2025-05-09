package parkingSpot

import (
	"parkingLot/vehicles"
	"sync"
)

func GetNewParkingSpot(id int, vehicleType vehicles.VehicleType) *ParkingSpot {
	return &ParkingSpot{
		id:          id,
		vehicleType: vehicleType,
	}
}

type ParkingSpot struct {
	id          int
	vehicle     vehicles.Vehicle
	vehicleType vehicles.VehicleType
}

func (p *ParkingSpot) GetId() int {
	return p.id
}

func (p *ParkingSpot) IsAvailable(mu *sync.Mutex) bool {
	mu.Lock()
	defer mu.Unlock()
	return p.vehicle == nil
}

func (p *ParkingSpot) GetType() vehicles.VehicleType {
	return p.vehicleType
}

func (p *ParkingSpot) ParkVehicle(v vehicles.Vehicle, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	p.vehicle = v
}

func (p *ParkingSpot) UnParkVehicle(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	p.vehicle = nil
}

func (p *ParkingSpot) GetVehicle() vehicles.Vehicle {
	return p.vehicle
}
