package main

import "fmt"

func main() {
	// Declaração explícita
	var nome string = "Dr. Go"

	// Declaração curta (mais comum), o compilador infere o tipo
	idade := 35
	experiente := true
	pi := 3.14159

	fmt.Printf("Nome: %s (tipo: %T)\n", nome, nome)
	fmt.Printf("Idade: %d (tipo: %T)\n", idade, idade)
	fmt.Printf("Experiente: %t (tipo: %T)\n", experiente, experiente)
	fmt.Printf("Pi: %f (tipo: %T)\n", pi, pi)
}
