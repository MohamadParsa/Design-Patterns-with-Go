package main

import (
	"fmt"

	"./creational/builder"
	"./creational/factory"

	"github.com/google/uuid"
)

func main() {
	createObjectWithBuilderPattern()
	createObjectWithFactoryPattern()

}

func createObjectWithFactoryPattern() {
	fmt.Println("-------Factory Results ----------")
	car, _ := factory.NewVehicle("car", "BMW", "black", "M240i xDrive Coupe")
	truck, _ := factory.NewVehicle("truck", "Volvo", "black", "FH16")
	car.Move()
	truck.Move()
	car.Stop()
	fmt.Printf("object: %v  \n", car.String())
	fmt.Printf("object: %v  \n", truck.String())
	fmt.Println("---------------------------------")
}

func createObjectWithBuilderPattern() {
	fmt.Println("--------Builder Results ---------")
	objectBuilder := builder.NewObjectBuilder()
	//to see error
	object, err := objectBuilder.Build()
	fmt.Printf("object: %v | err: %v \n", object, err)

	id, _ := uuid.NewUUID()
	objectBuilder.SetID(id)
	objectBuilder.SetName("Parsa")
	objectBuilder.SetCount(10)

	object, err = objectBuilder.Build()
	fmt.Printf("object: %+v | err: %v \n", object, err)
	fmt.Println("---------------------------------")
}
