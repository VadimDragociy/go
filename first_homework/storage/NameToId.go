package storage

type NameToId struct {
	Name_to_id map[string]string
}

func NewNameToId() *NameToId {
	return &NameToId{Name_to_id: make(map[string]string)}
}
