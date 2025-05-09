package vehicles

func GetNewTruck(license string) Vehicle {
	return &BaseVehicle{
		license:     license,
		vehicleType: Truck,
	}
}
