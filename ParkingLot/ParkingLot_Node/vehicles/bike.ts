import { Vehicle, VehicleType } from "./vehicle.ts";

export default class Bike extends Vehicle {
  constructor(license: string) {
    super(license, VehicleType.bike);
  }
}
