import { Vehicle, VehicleType } from "./vehicle.ts";

export default class Truck extends Vehicle {
  constructor(license: string) {
    super(license, VehicleType.truck);
  }
}
