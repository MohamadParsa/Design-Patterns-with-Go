package abstract_factory

type ExpensiveHotel struct {
	address   string
	equipment string
}

func (expensiveHotel *ExpensiveHotel) SetAddress(address string) {
	expensiveHotel.address = address
}
func (expensiveHotel *ExpensiveHotel) SetEquipment(equipment string) {
	expensiveHotel.equipment = equipment
}
func (expensiveHotel *ExpensiveHotel) GetAddress() string {
	return expensiveHotel.address
}
func (expensiveHotel *ExpensiveHotel) GetEquipment() string {
	return expensiveHotel.equipment
}
