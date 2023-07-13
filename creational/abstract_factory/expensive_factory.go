package abstract_factory

import (
	exp "github.com/MohamadParsa/Design-Patterns-with-Go/creational/abstract_factory/expensive"
)

type Expensive struct {
}

func (expensive *Expensive) CreateHotel() Hotel {
	expensiveHotel := &exp.ExpensiveHotel{}
	expensiveHotel.SetAddress("Meadown")
	expensiveHotel.SetEquipment("sauna room")
	return expensiveHotel
}
func (expensive *Expensive) CreateTransport() Transport {
	expensiveTransport := &exp.ExpensiveTransport{}
	expensiveTransport.SetTransportType("taxi")
	expensiveTransport.SetService("load handling")
	return expensiveTransport
}
