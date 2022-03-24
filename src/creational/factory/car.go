package factory

type Car struct {
	Vehicle
}

func (car *Car) GetStatus() string {
	return car.Vehicle.status
}
func (car *Car) String() string {
	return "car " + car.Vehicle.brand + " " + car.Vehicle.color + " " + car.Vehicle.model + " " + car.Vehicle.status
}

func newCar(brand, color, model string) *Car {
	return &Car{Vehicle: Vehicle{brand: brand, color: color, model: model}}
}
