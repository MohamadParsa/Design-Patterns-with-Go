package factory

//Truck Struct is a type of vehicle so we use composition to inject vehicle attributes into truck.
type Truck struct {
	Vehicle
}

//GetStatus returns Moving status of the truck.
func (truck *Truck) GetMovingStatus() string {
	return truck.Vehicle.status
}

//String returns information of the truck as a text.
func (truck *Truck) String() string {
	return "truck " + truck.Vehicle.brand + " " + truck.Vehicle.color + " " + truck.Vehicle.model + " " + truck.Vehicle.status
}
func newTruck(brand, color, model string) *Truck {
	return &Truck{Vehicle: Vehicle{brand: brand, color: color, model: model}}
}
