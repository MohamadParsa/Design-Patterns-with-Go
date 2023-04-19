package iterator

//library struct is data collection.
type Library struct {
	books []*Book
}

//AddBook function add a book to the library.
func (library *Library) AddBook(book Book) {
	library.books = append(library.books, &book)
}

//GetCollection function returns a iterate able collection.
func (library *Library) GetCollection() IteratorInterface[Book] {
	return &Books{books: library.books, current: -1}
}
