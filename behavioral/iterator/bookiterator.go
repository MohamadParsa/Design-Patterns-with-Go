package iterator

type Books struct {
	current int
	books   []*Book
}
type Book struct {
	BookName string
	Author   string
}

func (books *Books) HasNext() bool {
	return books.current < len(books.books)-1
}

func (books *Books) HasPrevious() bool {
	return books.current > 0
}
func (books *Books) Next() *Book {
	if books.HasNext() {
		books.current++
	}
	return books.books[books.current]
}
func (books *Books) Previous() *Book {
	if books.HasPrevious() {
		books.current--
	}
	return books.books[books.current]
}
