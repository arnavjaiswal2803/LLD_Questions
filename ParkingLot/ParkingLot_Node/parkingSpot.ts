import { Vehicle, VehicleType } from "./vehicles/vehicle.ts";

export default class ParkingSpot {
  _spotNumber: number;
  _vehicle: Vehicle | null;
  _vehicleType: VehicleType;

  constructor(spotNumber: number, vehicleType: VehicleType) {
    this._spotNumber = spotNumber;
    this._vehicleType = vehicleType;
    this._vehicle = null;
  }

  get spotNumber() {
    return this._spotNumber;
  }

  get vehicle() {
    return this._vehicle;
  }

  get vehicleType() {
    return this._vehicleType;
  }

  set vehicle(v: Vehicle | null) {
    this._vehicle = v;
  }

  isAvailable() {
    return this.vehicle === null;
  }

  parkVehicle(v: Vehicle) {
    if (this.isAvailable() && this.vehicleType === v.vehicleType) {
      this.vehicle = v;
      return true;
    }
    return false;
  }

  unParkVehicle() {
    if (!this.isAvailable()) {
      this.vehicle = null;
      return true;
    }
    return false;
  }
}
