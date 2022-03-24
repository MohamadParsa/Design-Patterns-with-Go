package factory

type VehicleInterface interface {
	SetBrand(string)
	GetBrand() string
	SetColor(string)
	GetColor() string
	SetModel(string)
	GetModel() string
	Move()
	Stop()
	GetStatus() string
	String() string
}

type Vehicle struct {
	brand  string
	color  string
	model  string
	status string
}

func (vehicle *Vehicle) GetBrand() string {
	return vehicle.brand
}
func (vehicle *Vehicle) SetBrand(brand string) {
	vehicle.brand = brand
}
func (vehicle *Vehicle) GetColor() string {
	return vehicle.color
}
func (vehicle *Vehicle) SetColor(color string) {
	vehicle.color = color
}
func (vehicle *Vehicle) GetModel() string {
	return vehicle.model
}
func (vehicle *Vehicle) SetModel(model string) {
	vehicle.model = model
}
func (vehicle *Vehicle) Move() {
	vehicle.status = "Moveing"
}
func (vehicle *Vehicle) Stop() {
	vehicle.status = "Stoped"
}
