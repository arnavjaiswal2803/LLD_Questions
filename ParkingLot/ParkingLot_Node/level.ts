import ParkingSpot from "./parkingSpot.ts";
import { Vehicle, VehicleType } from "./vehicles/vehicle.ts";

export default class Level {
  _floorNumber: number;
  _parkingSpots: ParkingSpot[];

  constructor(floorNumber: number, spotsCount: number) {
    this._floorNumber = floorNumber;
    this._parkingSpots = new Array(spotsCount);

    const bikeSpots = Math.floor(spotsCount * 0.6);
    for (let i = 0; i < bikeSpots; i++) {
      this._parkingSpots[i] = new ParkingSpot(i, VehicleType.bike);
    }

    const carSpots = Math.floor(spotsCount * 0.3);
    for (let i = bikeSpots; i < bikeSpots + carSpots; i++) {
      this._parkingSpots[i] = new ParkingSpot(i, VehicleType.car);
    }

    for (let i = bikeSpots + carSpots; i < spotsCount; i++) {
      this._parkingSpots[i] = new ParkingSpot(i, VehicleType.truck);
    }
  }

  get floorNumber() {
    return this._floorNumber;
  }

  get parkingSpots() {
    return this._parkingSpots;
  }

  parkVehicle(v: Vehicle) {
    const n = this._parkingSpots.length;

    for (let i = 0; i < n; i++) {
      const parkingSpot = this._parkingSpots[i];

      if (
        parkingSpot.isAvailable() &&
        parkingSpot.vehicleType === v.vehicleType
      ) {
        parkingSpot.parkVehicle(v);
        return true;
      }
    }
    return false;
  }

  unParkVehicle(v: Vehicle) {
    const n = this._parkingSpots.length;

    for (let i = 0; i < n; i++) {
      const parkingSpot = this._parkingSpots[i];

      if (parkingSpot.vehicle === v) {
        parkingSpot.unParkVehicle();
        return true;
      }
    }
    return false;
  }

  checkAvailability() {
    const n = this._parkingSpots.length;
    let count = 0;

    for (let i = 0; i < n; i++) {
      const parkingSpot = this._parkingSpots[i];

      if (parkingSpot.isAvailable()) {
        count++;
      }
    }

    return count;
  }
}
