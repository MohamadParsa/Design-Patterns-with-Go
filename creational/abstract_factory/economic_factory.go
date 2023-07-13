package abstract_factory

import (
	eco "github.com/MohamadParsa/Design-Patterns-with-Go/creational/abstract_factory/economic"
)

type Economic struct {
}

func (economic *Economic) CreateHotel() Hotel {
	economicHotel := &eco.EconomicHotel{}
	economicHotel.SetAddress("Fanwood")
	return economicHotel
}
func (economic *Economic) CreateTransport() Transport {
	economicTransport := &eco.EconomicTransport{}
	economicTransport.SetTransportType("bus")
	return economicTransport
}
