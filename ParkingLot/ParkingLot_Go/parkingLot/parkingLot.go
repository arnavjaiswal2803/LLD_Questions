package parkingLot

import (
	"parkingLot/level"
	"parkingLot/vehicles"
)

var parkingLot *ParkingLot

func GetNewParkingLot() *ParkingLot {
	if parkingLot == nil {
		parkingLot = &ParkingLot{
			floors: 0,
			levels: make([]level.Level, 0),
		}
	}
	return parkingLot
}

type ParkingLot struct {
	floors int
	levels []level.Level
}

func (p *ParkingLot) AddLevel(spots int) {
	level := level.GetNewLevel(p.floors, spots)
	p.levels = append(p.levels, *level)
	p.floors++
}

func (p *ParkingLot) ParkVehicle(v vehicles.Vehicle) bool {
	for _, level := range p.levels {
		if level.ParkVehicle(v) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) UnParkVehicle(v vehicles.Vehicle) bool {
	for _, level := range p.levels {
		if level.UnParkVehicle(v) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) CheckAvailability() int {
	count := 0
	for _, level := range p.levels {
		count += level.CheckAvailability()
	}
	return count
}
