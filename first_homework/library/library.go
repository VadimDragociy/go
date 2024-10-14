package library

import (
	"github.com/VadimDragociy/go/book"
	"github.com/VadimDragociy/go/storage"

	"github.com/google/uuid"
)

type Library interface {
	Search(name string) (book.Book, bool)
	AddBook(book book.Book) bool         //если раскомментировать, то LibraryStock не реализует Library, потому что в AddBook pointer receiver (что логично, но я не понимаю как это обойти)
	setGenId(generator func() uuid.UUID) //то же самое
}

type LibraryStock struct {
	Storage storage.Storage
}

func NewLibraryStock(storage storage.Storage) *LibraryStock {
	return &LibraryStock{Storage: storage}
}

func (lib LibraryStock) Search(name string) (book.Book, error) {
	book, err := lib.Storage.Search(name)

	return book, err
}

func (lib *LibraryStock) AddBook(book book.Book) error {
	err := lib.Storage.AddBook(book)

	return err
}

func (lib *LibraryStock) SetGenId(generator func() uuid.UUID) {
	lib.Storage.ClearAndRegenId(generator)

}
