package storage

import (
	"errors"

	"github.com/VadimDragociy/go/book"
	"github.com/google/uuid"
)

type StorageMap struct {
	Vault_map map[string]book.Book

	Name_to_id NameToId

	Generator_new bool
	Generator     func() uuid.UUID
}

func NewStorageMap() *StorageMap {
	return &StorageMap{Vault_map: make(map[string]book.Book), Name_to_id: *NewNameToId(), Generator_new: false}
}

func (lib StorageMap) GetIdFromName(name string) (string, bool) {
	id, ok := lib.Name_to_id.Name_to_id[name]
	return id, ok
}

func (searcher StorageMap) GetBookByid(id string) book.Book {
	book := searcher.Vault_map[id]
	return book
}

func (lib *StorageMap) AddBook(book book.Book) error {
	_, ok := lib.Name_to_id.Name_to_id[book.Name]
	len := len(lib.Name_to_id.Name_to_id)
	if !ok || len == 0 {
		id := lib.genId()
		book.SetId(id)
		lib.Vault_map[book.Id] = book
		lib.Name_to_id.Name_to_id[book.Name] = book.Id
		return nil
	}
	return errors.New("book already exists")
}

func (lib StorageMap) genId() string {
	var id = ""
	if lib.Generator_new {
		id = lib.Generator().String()
	} else {
		id = uuid.New().String()
	}
	return id
}

func (lib *StorageMap) ClearAndRegenId(generator func() uuid.UUID) {
	lib.Generator_new = true
	lib.Generator = generator

	for k := range lib.Vault_map {
		new_id := lib.Generator()
		book := lib.GetBookByid(k)
		lib.Name_to_id.Name_to_id[book.Name] = new_id.String()

		book.SetId(new_id.String())
		delete(lib.Vault_map, k)
		lib.Vault_map[book.Id] = book
	}

}

func (lib StorageMap) Search(name string) (book.Book, error) {
	id, ok := lib.GetIdFromName(name)
	if !ok {
		return book.Book{}, errors.New("book doesn`t exist")
	}
	book := lib.GetBookByid(id)
	return book, nil
}
