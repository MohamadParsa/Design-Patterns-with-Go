package iterator

//CollactionInterface determines what functions the data collection should to have.
type CollactionInterface[itemType any] interface {
	GetCollection() *IteratorInterface[itemType]
}
