package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Book struct {
	Name string
	id   string
}

type Searcher struct {
	lib *Library
}

type Library struct {
	vault_slice []Book
	vault_map   map[string]Book
	name_to_id  map[string]string
	searcher    *Searcher

	type_o_vault  bool
	generator     func() uuid.UUID
	generator_new bool
}

func (lib Library) get_id_from_name(name string) (string, bool) {
	id, ok := lib.name_to_id[name]
	return id, ok
}

//finished

func (searcher Searcher) Search(id string) Book {
	if searcher.lib.type_o_vault {
		book := searcher.lib.vault_map[id]
		return book
	}
	vault := searcher.lib.vault_slice
	size := len(searcher.lib.vault_slice)
	for i := 0; i < size; i++ {
		if vault[i].id == id {
			return vault[i]
		}
	}
	return Book{} //костыль
}

//finished

func (lib Library) Search(name string) (Book, bool) {
	id, ok := lib.get_id_from_name(name)
	if ok {
		book := lib.searcher.Search(id)
		return book, ok
	}
	return Book{}, ok
}

//finished

func (lib *Library) Switch_type_o_vault(newtype bool) {
	lib.type_o_vault = newtype
}

//finished

func (lib *Library) Add_a_book(book Book) bool {
	if len(lib.vault_map) > 0 {
		_, ok := lib.Search(book.Name)
		if ok {
			return !ok
		}
	}

	id := lib.gen_id()
	book.set_id(id)

	lib.name_to_id[book.Name] = book.id
	lib.vault_map[book.id] = book
	lib.vault_slice = append(lib.vault_slice, book)

	return true
}

//finished

func (b *Book) set_id(id string) Book {
	b.id = id
	return *b
}

//finished

func (lib *Library) gen_id() string {
	var id = ""
	if lib.generator_new {

		id = lib.generator().String()
	} else {
		id = uuid.New().String()
	}

	return id
}

//finished

func (lib *Library) set_gen_id(generator func() uuid.UUID) {
	lib.generator_new = true
	lib.generator = generator

	for k := range lib.vault_map {
		delete(lib.vault_map, k)
	}

	for i, book := range lib.vault_slice {
		new_id := lib.generator()
		lib.name_to_id[book.Name] = new_id.String()

		lib.vault_slice[i].set_id(new_id.String())
		book.set_id(new_id.String())

		lib.vault_map[book.id] = book
	}

}

//finished

func main() {
	vault_slice := []Book{}
	vault_map := map[string]Book{}
	name_to_id := map[string]string{}
	library := Library{}
	searcher := Searcher{}

	library = Library{
		vault_map: vault_map, name_to_id: name_to_id, vault_slice: vault_slice, searcher: &searcher,
	}
	searcher = Searcher{
		lib: &library,
	}
	books_arrive1 := []Book{
		Book{Name: "Cool scenarios for drama"}, Book{Name: "Cooking with spoons"}, Book{Name: "The BusError"}, Book{Name: "Taburetka"},
	}
	books_arrive2 := []Book{
		Book{Name: "Being wrong"}, Book{Name: "The Frog"}, Book{Name: "The Good, the Bad and the Segmentation fauld"}, Book{Name: "Salute"},
	}

	for i := 0; i < len(books_arrive1); i++ {
		library.Add_a_book(books_arrive1[i])
	}

	book1, _ := library.Search("Taburetka")
	book2, _ := library.Search("Cooking with spoons")

	fmt.Println("I hate ", book1.Name, book1.id)
	fmt.Println("I hate ", book2.Name, book2.id)

	library.set_gen_id(uuid.New) // использую ту же самую функцию-генератор айди, но они должны перегенерироваться
	book3, _ := library.Search("Taburetka")
	book4, _ := library.Search("The BusError")

	fmt.Println("I hate ", book3.Name, book3.id)
	fmt.Println("I hate ", book4.Name, book4.id)

	library.Switch_type_o_vault(true)

	for i := 0; i < len(books_arrive2); i++ {
		library.Add_a_book(books_arrive2[i])
	}
	// book5, _ := library.Search("Salute")
	book6, _ := library.Search("The Good, the Bad and the Segmentation fauld")

	// fmt.Println("I hate ", book5.Name, book5.id)
	fmt.Println("I hate ", book6.Name, book6.id)

}
