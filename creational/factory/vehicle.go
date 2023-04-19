package factory

//VehicleInterface determines what functions our structs have to have.
type VehicleInterface interface {
	SetBrand(string)
	GetBrand() string
	SetColor(string)
	GetColor() string
	SetModel(string)
	GetModel() string
	Move()
	Stop()
	GetMovingStatus() string
	String() string
}

//Vehicle contains common properties of any vehicle type.
type Vehicle struct {
	brand  string
	color  string
	model  string
	status string
}

//GetBrand return brand of the vehicle.
func (vehicle *Vehicle) GetBrand() string {
	return vehicle.brand
}

//SetModel set brand of the vehicle.
func (vehicle *Vehicle) SetBrand(brand string) {
	vehicle.brand = brand
}

//GetBrand return color of the vehicle.
func (vehicle *Vehicle) GetColor() string {
	return vehicle.color
}

//SetModel set color of the vehicle.
func (vehicle *Vehicle) SetColor(color string) {
	vehicle.color = color
}

//GetBrand return model of the vehicle.
func (vehicle *Vehicle) GetModel() string {
	return vehicle.model
}

//SetModel set model of the vehicle.
func (vehicle *Vehicle) SetModel(model string) {
	vehicle.model = model
}

//Move change the vehicle status to Moving.
func (vehicle *Vehicle) Move() {
	vehicle.status = "Moveing"
}

//Stop change the vehicle status to Stopped.
func (vehicle *Vehicle) Stop() {
	vehicle.status = "Stopped"
}
