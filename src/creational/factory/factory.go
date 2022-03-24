//Package factory is an implementation of factory pattern in Go
package factory

import "errors"

/*Factory method is a creational design pattern, i.e., related to object creation. In Factory pattern,
we create objects without exposing the creation logic to the client and the client uses the same
common interface to create a new type of object.
The idea is to use a static member-function (static factory method) that creates & returns instances,
hiding the details of class modules from the user.
A factory pattern is one of the core design principles to create an object, allowing clients to
create objects of a library(explained below) in a way such that it doesn’t have tight coupling with
the class hierarchy of the library.
refrence: https://www.geeksforgeeks.org/design-patterns-set-2-factory-method/
*/

//all returned errors defined to use in test.
var (
	errorVehicleTypeIsInvalid = errors.New("vehicleType is invalid")
)

//factory method to return object.
func NewVehicle(vehicleType, brand, color, model string) (VehicleInterface, error) {
	if vehicleType == "car" {
		return newCar(brand, color, model), nil
	}
	if vehicleType == "truck" {
		return newTruck(brand, color, model), nil
	}

	return newCar("", "", ""), errorVehicleTypeIsInvalid
}
