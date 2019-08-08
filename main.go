package main

type Item struct {
	ID   string
	Name string
}

func GetItem(id, name string) *Item {
	return &Item{
		ID:   id,
		Name: name,
	}
}
