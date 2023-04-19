package adapter

import "fmt"

//PlasticToyDuck is a struct to implement plastic toy duck actions.
type PlasticToyDuck struct {
	squeaking bool
}

//Squeak function tells the plastic toy duck its squeak time.
func (plasticToyDuck *PlasticToyDuck) Squeak() {
	plasticToyDuck.squeaking = true
	fmt.Println("plastic toy duck is squeaking")
}
