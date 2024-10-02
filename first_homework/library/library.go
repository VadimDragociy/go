package library

import (
	"errors"
	"sample-app/first_homework/book"
	"sample-app/first_homework/storage"

	"github.com/google/uuid"
)

type NameToId struct {
	Name_to_id map[string]string
}

type Library interface {
	Search(name string) (book.Book, bool)
	// AddBook(book Book) bool //если раскомментировать, то LibraryStock не реализует Library, потому что в AddBook pointer receiver (что логично, но я не понимаю как это обойти)
	getIdFromName(name string) (string, error)
	getBookByid(id string) book.Book
	genId() string
	// setGenId(generator func() uuid.UUID) //то же самое
}

// type Searcher interface {
// 	Search(id string) (Book, bool)
// 	SwitchTypeOfVault(newtype bool)
// }

type LibraryStock struct {
	Vault_map    storage.StorageMap
	Vault_slice  storage.StorageSlice
	Type_o_vault bool
	Name_to_id   NameToId

	Generator_new bool
	Generator     func() uuid.UUID
}

func (lib LibraryStock) GetIdFromName(name string) (string, bool) {
	id, ok := lib.Name_to_id.Name_to_id[name]
	return id, ok
}

func (searcher LibraryStock) GetBookByid(id string) book.Book { //здесь не требуется дополнительной проверки, потому что о наличии/ отсутствии книги мы знаем уже на строке 92

	if searcher.Type_o_vault {
		book := searcher.Vault_map.Vault_map[id]
		return book
	}
	vault := searcher.Vault_slice.Vault_slice
	size := len(searcher.Vault_slice.Vault_slice)
	for i := 0; i < size; i++ {
		if vault[i].Id == id {
			return vault[i]
		}
	}
	return book.Book{}
}

//finished

// func (searcher StorageSlice) Search(id string) Book {
// 	book := Book{}
// 	vault := searcher.vault_slice
// 	size := len(searcher.vault_slice)
// 	for i := 0; i < size; i++ {
// 		if vault[i].id == id {
// 			book = vault[i]
// 		}
// 	}
// 	return book

// }
//finished

func (lib LibraryStock) Search(name string) (book.Book, error) {
	id, ok := lib.GetIdFromName(name)
	if ok {
		book := lib.GetBookByid(id)
		return book, nil
	}
	return book.Book{}, errors.New("no book with such name was found")
}

//finished

func (lib *LibraryStock) SwitchTypeOfVault(newtype bool) {
	lib.Type_o_vault = newtype
}

//finished

func (lib *LibraryStock) AddBook(book book.Book) error {
	if len(lib.Vault_map.Vault_map) > 0 {
		_, err := lib.Search(book.Name)
		if err == nil {
			return errors.New("Book already exists")
		}
	}

	id := lib.genId()
	book.SetId(id)

	lib.Name_to_id.Name_to_id[book.Name] = book.Id
	lib.Vault_map.Vault_map[book.Id] = book
	lib.Vault_slice.Vault_slice = append(lib.Vault_slice.Vault_slice, book)

	return nil
}

//finished

func (lib LibraryStock) genId() string {
	var id = ""
	if lib.Generator_new {

		id = lib.Generator().String()
	} else {
		id = uuid.New().String()
	}

	return id
}

//finished

func (lib *LibraryStock) SetGenId(generator func() uuid.UUID) {
	lib.Generator_new = true
	lib.Generator = generator

	for k := range lib.Vault_map.Vault_map {
		delete(lib.Vault_map.Vault_map, k)
	}

	for i, book := range lib.Vault_slice.Vault_slice {
		new_id := lib.Generator()
		lib.Name_to_id.Name_to_id[book.Name] = new_id.String()

		lib.Vault_slice.Vault_slice[i].SetId(new_id.String())
		book.SetId(new_id.String())

		lib.Vault_map.Vault_map[book.Id] = book
	}

}
