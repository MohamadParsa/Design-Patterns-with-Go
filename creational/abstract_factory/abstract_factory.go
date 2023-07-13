package abstract_factory

import (
	"errors"
)

/*
Abstract Factory design pattern is one of the Creational pattern. Abstract Factory pattern is almost similar to Factory Pattern and is considered
as another layer of abstraction over factory pattern. Abstract Factory patterns work around a super-factory which creates other factories.
Abstract factory pattern implementation provides us with a framework that allows us to create objects that follow a general pattern.
So at runtime, the abstract factory is coupled with any desired concrete factory which can create objects of the desired type.
reference: https://www.geeksforgeeks.org/abstract-factory-pattern/
*/

type ReservationFactory interface {
	CreateHotel() Hotel
	CreateTransport() Transport
}

func GetReservationFactory(reservationType string) (ReservationFactory, error) {
	switch reservationType {
	case "economic":
		return &Economic{}, nil
	case "expensive":
		return &Expensive{}, nil
	}
	return nil, errors.New("wrong type for reservation")
}
