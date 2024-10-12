package storage

import (
	"github.com/VadimDragociy/go/book"

	"github.com/google/uuid"
)

type Storage interface {
	GetBookByid(id string) book.Book
	AddBook(book book.Book) error
	ClearAndRegenId(generator func() uuid.UUID)
	GetIdFromName(name string) (string, bool)
	Search(name string) (book.Book, error)
	genId() string
}
