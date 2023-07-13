// Package adapter is an implementation of adapter pattern in Go
package adapter

/*This pattern is easy to understand as the real world is full of adapters.
For example consider a USB to Ethernet adapter. We need this when we have an Ethernet interface on
one end and USB on the other. Since they are incompatible with each other. we use an adapter that
converts one to other. This example is pretty analogous to Object Oriented Adapters. In design,
adapters are used when we have a class (Client) expecting some type of object and we have an object
(Adaptee) offering the same features but exposing a different interface.
reference: https://www.geeksforgeeks.org/adapter-pattern/?ref=lbp
*/

// BirdAdapter struct adapte the Sparrow struct to ToyDuck
type BirdAdapter struct {
	Sparrow
}

// Squeak function tells the bird its squeak time.
func (birdAdapter *BirdAdapter) Squeak() {
	//Squeak is a required function according to the ToyDuck interface so we call the equal method in sparrow struct.
	birdAdapter.Sparrow.MakeSound()
}
