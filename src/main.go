package main

import (
	"fmt"

	"strconv"

	"./behavioral/observer"
	"./creational/builder"
	"./creational/factory"
	"./creational/singleton"
	"./structural/adapter"
	"./structural/facade"

	"github.com/google/uuid"
)

func main() {
	//calling Creational Patterns examples.
	createObjectWithBuilderPattern()
	createObjectWithFactoryPattern()
	createObjectInstanceWithSingletonPattern()
	//calling Structural Patterns examples.
	adaptSparrowAndToyDuckWithAdapterPattern()
	getMenuFromMultiResturantWithFacadePattern()
	//calling Structural Patterns examples.
	updateAllClientsWithObserverPattern()
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
func createObjectInstanceWithSingletonPattern() {
	channel := make(chan string, 10)
	fmt.Println("-------Singleton Results --------")
	i := 0
	for i = 0; i < 5; i++ {
		go func(i int) {
			singletonObject := singleton.GetInstance()
			singletonObject.GenerateID()
			channel <- " Loop Number: " + strconv.Itoa(i) + " , ID: " + singletonObject.GetID().String()
		}(i)
	}
	for {
		select {
		case message, status := <-channel:
			if status {
				fmt.Println(message)
				i--
				if i == 0 {
					close(channel)
				}
			} else {
				fmt.Println("---------------------------------")
				return
			}

		}
	}
}
func adaptSparrowAndToyDuckWithAdapterPattern() {
	fmt.Println("--------Adapter Results ---------")
	sparrow := adapter.Sparrow{}
	toyDuck := adapter.PlasticToyDuck{}

	// Wrap a bird in a birdAdapter so that it
	// behaves like toy duck
	birdAdapter := adapter.BirdAdapter{}

	fmt.Println("Sparrow...")
	sparrow.Fly()
	sparrow.MakeSound()
	fmt.Println("ToyDuck...")
	toyDuck.Squeak()

	// toy duck behaving like a bird
	fmt.Println("BirdAdapter...")
	birdAdapter.Squeak()
	fmt.Println("---------------------------------")
}
func getMenuFromMultiResturantWithFacadePattern() {
	fmt.Println("--------Facade Results ---------")

	fmt.Println("Vegetarian Restaurant Menu: ", facade.GetVegMenu())
	fmt.Println("Non Vegetarian Restaurant Menu: ", facade.GetNonVegMenu())
	fmt.Println("---------------------------------")
}
func updateAllClientsWithObserverPattern() {
	fmt.Println("--------Observer Results ---------")
	subject := observer.Subject{Observers: make(map[uuid.UUID]*observer.Observer)}
	observer1 := &observer.Observer{}
	observer1.SetName("Tom")
	subject.RegisterOnObserveList(observer1)
	observer2 := &observer.Observer{}
	observer2.SetName("Niki")
	subject.RegisterOnObserveList(observer2)

	milk := subject.NewProduct("Milk", 2)
	meat := subject.NewProduct("meat", 19)

	subject.Updateall(milk)
	subject.Updateall(meat)
	milk.Price = 1
	subject.Updateall(milk)
	subject.RemoveFromObserveList(observer2)
	fmt.Println("---------------------------------")
}
