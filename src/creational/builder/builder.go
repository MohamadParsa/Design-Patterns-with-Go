//Package builder is an implementation of builder pattern in Go
package builder

import (
	"errors"

	"github.com/google/uuid"
)

/*
Builder pattern aims to “Separate the construction of a complex object from its representation
so that the same construction process can create different representations.” It is used to
construct a complex object step by step and the final step will return the object. The process
of constructing an object should be generic so that it can be used to create different
representations of the same object.
refrence: https://www.geeksforgeeks.org/builder-design-pattern/
*/

//object is our complex object that we want to create.
type ObjectBuilder struct {
	id          uuid.UUID
	name        string
	count       int
	description string
}

//all returned errors defined to use in test.
var (
	errorIdIsEmpty      = errors.New("id can not be empty")
	errorNameIsEmpty    = errors.New("name can not be empty")
	errorCountIsInvalid = errors.New("count can not be less than 0")
)

//NewObjectBuilder return a builder to create a new object.
func NewObjectBuilder() *ObjectBuilder {
	return &ObjectBuilder{}
}

//Build return an instance of the object and if there is any error returns an error
func (b *ObjectBuilder) Build() (*Object, error) {
	if err := b.checkRequiredItem(); err != nil {
		return nil, err
	}
	return &Object{
		id:          b.id,
		name:        b.name,
		count:       b.count,
		description: b.description,
	}, nil
}

//SetID sets the ID of the object
func (b *ObjectBuilder) SetID(id uuid.UUID) {
	b.id = id
}

//SetName sets the Name of the object
func (b *ObjectBuilder) SetName(name string) {
	b.name = name
}

//SetCount sets the Cont of the object
func (b *ObjectBuilder) SetCount(count int) {
	b.count = count
}

//SetDescription sets the Description of the object
func (b *ObjectBuilder) SetDescription(description string) {
	b.description = description
}

func (b *ObjectBuilder) checkRequiredItem() error {
	if b.id == uuid.Nil {
		return errorIdIsEmpty
	}
	if b.name == "" {
		return errorNameIsEmpty
	}
	if b.count < 0 {
		return errorCountIsInvalid
	}
	return nil
}
