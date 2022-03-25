package adapter

import "fmt"

type Sparrow struct {
	flying      bool
	makingSound bool
}

func (sparrow *Sparrow) Fly() {
	sparrow.flying = true
	sparrow.makingSound = false
	fmt.Println("sparrow is flying")
}
func (sparrow *Sparrow) MakeSound() {
	sparrow.flying = false
	sparrow.makingSound = true
	fmt.Println("sparrow is making sound")

}
