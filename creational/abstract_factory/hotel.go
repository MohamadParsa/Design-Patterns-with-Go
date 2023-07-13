package abstract_factory

type HotelFactory interface {
}

type Hotel interface {
	GetAddress() string
	GetEquipment() string
}
