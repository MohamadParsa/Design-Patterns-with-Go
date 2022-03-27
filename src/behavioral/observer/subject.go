package observer

import (
	"fmt"

	"github.com/google/uuid"
)

type SubjectInterface interface {
	RegisterOnObserveList(*Observer)
	RemoveFromObserveList(*Observer)
	Updateall(ProductPrice)
}

type Subject struct {
	Observers map[uuid.UUID]*Observer
}

func (subject *Subject) RegisterOnObserveList(observer *Observer) {
	fmt.Println(observer.name + " registered.")
	subject.Observers[observer.id] = observer
}
func (subject *Subject) RemoveFromObserveList(observer *Observer) {
	fmt.Println(observer.name + " removed.")
	delete(subject.Observers, observer.id)
}
func (subject *Subject) Updateall(productPrice ProductPrice) {
	for _, observer := range subject.Observers {
		observer.onUpdate(productPrice)
	}
}
func (subject *Subject) NewProduct(name string, price int) ProductPrice {
	id, _ := uuid.NewUUID()
	return ProductPrice{ID: id, Name: name, Price: price}
}
