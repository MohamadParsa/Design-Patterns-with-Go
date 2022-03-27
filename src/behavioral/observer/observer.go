//Package observer is an implementation of observer pattern in Go
package observer

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

/*The observer pattern is a software design pattern in which an object, named the subject, maintains
a list of its dependents, called observers, and notifies them automatically of any state changes,
usually by calling one of their methods.
It is mainly used for implementing distributed event handling systems, in "event driven" software.
In those systems, the subject is usually named a "stream of events" or "stream source of events",
while the observers are called "sinks of events". The stream nomenclature alludes to a physical setup
where the observers are physically separated and have no control over the emitted events from the
subject/stream-source. This pattern then perfectly suits any process where data arrives from some
input that is not available to the CPU at startup, but instead arrives "at random" (HTTP requests,
GPIO data, user input from keyboard/mouse/..., distributed databases and blockchains, ...). Most
modern programming-languages comprise built-in "event" constructs implementing the observer-pattern
components. While not mandatory, most 'observers' implementations would use background threads
listening for subject-events and other support mechanisms provided by the kernel (Linux epoll, ...).
refrence: https://en.wikipedia.org/wiki/Observer_pattern
*/

type ObserverInterface interface {
	onUpdate(ProductPrice)
}
type Observer struct {
	id   uuid.UUID
	name string
	data map[uuid.UUID]ProductPrice
}
type ProductPrice struct {
	ID    uuid.UUID
	Name  string
	Price int
}

func (observer *Observer) SetName(name string) {
	observer.name = name
	observer.id, _ = uuid.NewUUID()
	observer.data = make(map[uuid.UUID]ProductPrice)
}

func (observer *Observer) onUpdate(productPrice ProductPrice) {
	product, exists := observer.data[productPrice.ID]
	observer.data[productPrice.ID] = productPrice

	//to see result
	if exists {
		fmt.Println(observer.name + " , got update price for " + productPrice.Name +
			" from " + strconv.Itoa(product.Price) + " to " + strconv.Itoa(productPrice.Price))
	} else {
		fmt.Println(observer.name + " , got " + productPrice.Name +
			" with price " + strconv.Itoa(productPrice.Price))
	}
}
