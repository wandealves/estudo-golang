package main

import "fmt"

type Usuario struct {
	ID    int
	Nome  string
	Email string
}

func main() {
	u := Usuario{ID: 1, Nome: "Alice", Email: "alice@email.com"}
	fmt.Printf("%+v\n", u) // O %+v imprime os nomes dos campos
}
