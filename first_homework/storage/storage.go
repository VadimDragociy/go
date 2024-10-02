package storage

import "sample-app/first_homework/book"

type StorageMap struct {
	Vault_map map[string]book.Book
}

type StorageSlice struct {
	Vault_slice []book.Book
}
