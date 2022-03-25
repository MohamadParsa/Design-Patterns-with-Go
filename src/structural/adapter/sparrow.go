package adapter

import "fmt"

//Sparrow is a struct to implement sparrow actions.
type Sparrow struct {
	flying      bool
	makingSound bool
}

//Fly function tells the bird it's flying time.
func (sparrow *Sparrow) Fly() {
	sparrow.flying = true
	sparrow.makingSound = false
	fmt.Println("sparrow is flying")
}

//MakeSound function tells the bird it's singing time.
func (sparrow *Sparrow) MakeSound() {
	sparrow.flying = false
	sparrow.makingSound = true
	fmt.Println("sparrow is making sound")

}
