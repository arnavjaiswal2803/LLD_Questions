import ParkingLot from "./parkingLot.ts";
import Truck from "./vehicles/truck.ts";
import Bike from "./vehicles/bike.ts";
import Car from "./vehicles/car.ts";

function demo() {
  const parkingLot = ParkingLot.instance;

  parkingLot.addLevel(10);

  const truck1 = new Truck("1234");
  const truck2 = new Truck("12345");
  const b1 = new Bike("231");
  const c1 = new Car("321");

  console.log(parkingLot.parkVehicle(truck1));
  console.log(parkingLot.parkVehicle(truck2));
  console.log(parkingLot.parkVehicle(b1));
  console.log(parkingLot.parkVehicle(c1));

  console.log(parkingLot.checkAvailability());

  console.log(parkingLot.unParkVehicle(truck2));
  console.log(parkingLot.unParkVehicle(truck1));

  console.log(parkingLot.checkAvailability());
}

demo();
