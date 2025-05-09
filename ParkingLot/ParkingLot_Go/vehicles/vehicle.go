package vehicles

type VehicleType int

const (
	Bike VehicleType = iota
	Car
	Truck
)

type Vehicle interface {
	GetType() VehicleType
	GetLicense() string
}

type BaseVehicle struct {
	license     string
	vehicleType VehicleType
}

func (v *BaseVehicle) GetType() VehicleType {
	return v.vehicleType
}

func (v *BaseVehicle) GetLicense() string {
	return v.license
}
