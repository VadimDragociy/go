package main

import (
	"fmt"
	"sample-app/first_homework/book"
	"sample-app/first_homework/library"
	"sample-app/first_homework/storage"

	"github.com/google/uuid"
)

//finished

func main() {
	vault_slice := []book.Book{}
	vault_map := map[string]book.Book{}
	name_to_id := map[string]string{}

	lib := library.LibraryStock{Vault_slice: storage.StorageSlice{vault_slice}, Vault_map: storage.StorageMap{vault_map}, Name_to_id: library.NameToId{name_to_id}}

	// var g Library = library

	books_arrive1 := []book.Book{
		{Name: "Cool scenarios for drama"}, {Name: "Cooking with spoons"}, {Name: "The BusError"}, {Name: "Taburetka"},
	}
	books_arrive2 := []book.Book{
		{Name: "Being wrong"}, {Name: "The Frog"}, {Name: "The Good, the Bad and the Segmentation fauld"}, {Name: "Salute"},
	}

	for i := 0; i < len(books_arrive1); i++ {
		lib.AddBook(books_arrive1[i])
	}

	book1, _ := lib.Search("Taburetka")
	book2, _ := lib.Search("Cooking with spoons")

	fmt.Println("I hate ", book1.Name, book1.Id)
	fmt.Println("I hate ", book2.Name, book2.Id)

	lib.SetGenId(uuid.New) // использую ту же самую функцию-генератор айди, но они должны перегенерироваться
	book3, _ := lib.Search("Taburetka")
	book4, _ := lib.Search("The BusError")

	fmt.Println("I hate ", book3.Name, book3.Id)
	fmt.Println("I hate ", book4.Name, book4.Id)

	lib.SwitchTypeOfVault(true)

	for i := 0; i < len(books_arrive2); i++ {
		lib.AddBook(books_arrive2[i])
	}
	book5, _ := lib.Search("Salute")
	book6, _ := lib.Search("The Good, the Bad and the Segmentation fauld")

	fmt.Println("I hate ", book5.Name, book5.Id)
	fmt.Println("I hate ", book6.Name, book6.Id)

	_, error1 := lib.Search("asdklsfdfbrjtnker")

	fmt.Println(error1)

}

// разбить проект по пакетам
// в библиотеке не должно быть хранения 					сделано
// пусть функция возвращает ощибку если книга не найдена	технически не сделано, но сделано по-другому
// CamelCase												сделано
// *Book 													сделано
// обработка ошибок											вроде бы сделано
// в Searcher только хранилище данных и все					??? переписал по-другому, потому что не понимаю, как имплементировать интерфейсы в код
