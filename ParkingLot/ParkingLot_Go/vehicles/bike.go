package vehicles

func GetNewBike(license string) Vehicle {
	return &BaseVehicle{
		license:     license,
		vehicleType: Bike,
	}
}
