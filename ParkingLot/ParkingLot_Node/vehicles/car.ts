import { Vehicle, VehicleType } from "./vehicle.ts";

export default class Car extends Vehicle {
  constructor(license: string) {
    super(license, VehicleType.car);
  }
}
