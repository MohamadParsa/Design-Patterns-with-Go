package builder

import "github.com/google/uuid"

//Object is our complex object
type Object struct {
	id          uuid.UUID
	name        string
	count       int
	description string
}
