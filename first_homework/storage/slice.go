package storage

import (
	"errors"

	"github.com/VadimDragociy/go/book"
	"github.com/google/uuid"
)

type StorageSlice struct {
	Vault_slice []book.Book

	Name_to_id NameToId

	Generator_new bool
	Generator     func() uuid.UUID
}

func NewStorageSlice() *StorageSlice {
	return &StorageSlice{Vault_slice: make([]book.Book, 0), Name_to_id: *NewNameToId(), Generator_new: false}
}

func (searcher StorageSlice) GetBookByid(id string) book.Book {
	vault := searcher.Vault_slice
	size := len(searcher.Vault_slice)
	for i := 0; i < size; i++ {
		if vault[i].Id == id {
			return vault[i]
		}
	}
	return book.Book{}
}

func (lib *StorageSlice) AddBook(book book.Book) error {
	_, ok := lib.Name_to_id.Name_to_id[book.Name]
	len := len(lib.Name_to_id.Name_to_id)
	if !ok || len == 0 {
		id := lib.genId()
		book.SetId(id)
		lib.Vault_slice = append(lib.Vault_slice, book)
		lib.Name_to_id.Name_to_id[book.Name] = book.Id
		return nil
	}
	return errors.New("book already exists")
}

func (lib StorageSlice) GetIdFromName(name string) (string, bool) {
	id, ok := lib.Name_to_id.Name_to_id[name]
	return id, ok
}

func (lib StorageSlice) genId() string {
	var id = ""
	if lib.Generator_new {
		id = lib.Generator().String()
	} else {
		id = uuid.New().String()
	}

	return id
}

func (lib *StorageSlice) ClearAndRegenId(generator func() uuid.UUID) {
	lib.Generator_new = true
	lib.Generator = generator

	for i, book := range lib.Vault_slice {
		new_id := lib.Generator()
		lib.Name_to_id.Name_to_id[book.Name] = new_id.String()

		lib.Vault_slice[i].SetId(new_id.String())
		book.SetId(new_id.String())
	}
}

func (lib StorageSlice) Search(name string) (book.Book, error) {
	id, ok := lib.GetIdFromName(name)
	if !ok {
		return book.Book{}, errors.New("book doesn`t exist")
	}
	book := lib.GetBookByid(id)
	return book, nil
}
