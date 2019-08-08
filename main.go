package main

type Item struct {
	ID    string
	Name  string
	token string
}

func GetItem(id, name string) *Item {
	token := "token"
	return &Item{
		ID:    id,
		Name:  name,
		token: token,
	}
}
