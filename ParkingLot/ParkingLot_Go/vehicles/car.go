package vehicles

func GetNewCar(license string) Vehicle {
	return &BaseVehicle{
		license:     license,
		vehicleType: Car,
	}
}
