package main

type Item struct {
	ID         string
	Name       string
	secretCode int64
}

func GetItem(id, name string) Item {
	return Item{
		ID:         id,
		Name:       name,
		secretCode: 12345,
	}
}
