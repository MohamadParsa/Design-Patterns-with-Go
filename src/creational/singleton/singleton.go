package singleton

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type object struct {
	id uuid.UUID
}

var instance *object
var one sync.Once

func (object *object) GetID() uuid.UUID {
	return object.id
}
func (object *object) GenerateID() {
	if object.id == uuid.Nil {
		object.id, _ = uuid.NewUUID()
	}
}

func GetInstance() *object {
	one.Do(func() {
		fmt.Println("Create!!!")
		instance = &object{}
	})
	return instance
}
