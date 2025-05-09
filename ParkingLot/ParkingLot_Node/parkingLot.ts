import Level from "./level.ts";
import { Vehicle } from "./vehicles/vehicle.ts";

export default class ParkingLot {
  static _instance: ParkingLot;

  _floorCount: number;
  _levels: Level[];

  private constructor() {
    this._floorCount = 0;
    this._levels = [];
  }

  static get instance() {
    if (!this._instance) {
      this._instance = new ParkingLot();
    }
    return this._instance;
  }

  get floorCount() {
    return this._floorCount;
  }

  get levels() {
    return this._levels;
  }

  addLevel(spotsCount: number) {
    this._levels.push(new Level(this.floorCount, spotsCount));
  }

  parkVehicle(v: Vehicle) {
    const n = this.levels.length;
    for (let i = 0; i < n; i++) {
      const level = this.levels[i];
      if (level.parkVehicle(v)) {
        return true;
      }
    }
    return false;
  }

  unParkVehicle(v: Vehicle) {
    const n = this.levels.length;
    for (let i = 0; i < n; i++) {
      const level = this.levels[i];
      if (level.unParkVehicle(v)) {
        return true;
      }
    }
    return false;
  }

  checkAvailability() {
    let count = 0;
    const n = this.levels.length;
    for (let i = 0; i < n; i++) {
      const level = this.levels[i];
      count += level.checkAvailability();
    }
    return count;
  }
}
