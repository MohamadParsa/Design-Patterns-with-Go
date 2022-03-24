package factory

type Truck struct {
	Vehicle
}

func (truck *Truck) GetStatus() string {
	return truck.Vehicle.status
}
func (truck *Truck) String() string {
	return "truck " + truck.Vehicle.brand + " " + truck.Vehicle.color + " " + truck.Vehicle.model + " " + truck.Vehicle.status
}
func newTruck(brand, color, model string) *Truck {
	return &Truck{Vehicle: Vehicle{brand: brand, color: color, model: model}}
}
