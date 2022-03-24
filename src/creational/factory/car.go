package factory

//Car Struct is a type of vehicle so we use composition to inject vehicle attributes into car.
type Car struct {
	Vehicle
}

//GetStatus returns Moving status of the car.
func (car *Car) GetMovingStatus() string {
	return car.Vehicle.status
}

//String returns information of the car as a text.
func (car *Car) String() string {
	return "car " + car.Vehicle.brand + " " + car.Vehicle.color + " " + car.Vehicle.model + " " + car.Vehicle.status
}

func newCar(brand, color, model string) *Car {
	return &Car{Vehicle: Vehicle{brand: brand, color: color, model: model}}
}
