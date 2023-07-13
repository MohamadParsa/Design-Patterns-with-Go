// Package iterator is an implementation of iterator pattern in Go
package iterator

/*the iterator pattern is a design pattern in which an iterator is used to traverse a container
and access the container's elements. The iterator pattern decouples algorithms from containers;
in some cases, algorithms are necessarily container-specific and thus cannot be decoupled.
reference: https://en.wikipedia.org/wiki/Iterator_pattern
*/

// IteratorInterface determines what functions the iterator should have.
type IteratorInterface[ItemType any] interface {
	HasNext() bool
	HasPrevious() bool
	Next() *ItemType
	Previous() *ItemType
}
