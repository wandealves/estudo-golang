package main

import "fmt"

func somar(a int, b int) (int, error) {
	if a+b > 100 {
		// A função fmt.Errorf cria um novo erro.
		return 0, fmt.Errorf("a soma não pode exceder 100")
	}
	// nil significa "sem erro".
	return a + b, nil
}

func main() {
	resultado, err := somar(50, 30)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
