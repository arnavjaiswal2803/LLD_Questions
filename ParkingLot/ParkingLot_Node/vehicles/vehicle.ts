export const VehicleType = {
  bike: "Bike",
  car: "Car",
  truck: "Truck",
};
export type VehicleType = (typeof VehicleType)[keyof typeof VehicleType];

export abstract class Vehicle {
  _license: string;
  _vehicleType: VehicleType;

  constructor(license: string, vehicleType: VehicleType) {
    this._license = license;
    this._vehicleType = vehicleType;
  }

  get license() {
    return this._license;
  }

  get vehicleType() {
    return this._vehicleType;
  }
}
