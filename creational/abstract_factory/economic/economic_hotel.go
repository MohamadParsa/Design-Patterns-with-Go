package economic

type EconomicHotel struct {
	address string
}

func (economicHotel *EconomicHotel) SetAddress(address string) {
	economicHotel.address = address
}
func (economicHotel *EconomicHotel) GetAddress() string {
	return economicHotel.address
}
func (economicHotel *EconomicHotel) GetEquipment() string {
	return "Nothing"
}
