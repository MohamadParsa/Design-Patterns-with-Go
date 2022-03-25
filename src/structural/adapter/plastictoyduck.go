package adapter

import "fmt"

type PlasticToyDuck struct {
	squeaking bool
}

func (plasticToyDuck *PlasticToyDuck) Squeak() {
	plasticToyDuck.squeaking = true
	fmt.Println("plastic toy duck is squeaking")

}
