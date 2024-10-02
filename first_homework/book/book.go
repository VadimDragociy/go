package book

type Book struct {
	Name string
	Id   string
}

func (b *Book) SetId(id string) *Book {
	b.Id = id
	return b
}
