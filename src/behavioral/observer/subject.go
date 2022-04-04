package observer

import (
	"fmt"

	"github.com/google/uuid"
)

type SubjectInterface interface {
	RegisterOnObserveList(*Observer)
	RemoveFromObserveList(*Observer)
	Updateall(Product)
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
func (subject *Subject) Updateall(product Product) {
	for _, observer := range subject.Observers {
		observer.onUpdate(product)
	}
}
func (subject *Subject) NewProduct(name string, price int) Product {
	id, _ := uuid.NewUUID()
	return Product{ID: id, Name: name, Price: price}
}
