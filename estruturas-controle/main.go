package main

import "fmt"

func main() {
	// Loop clássico
	for i := 0; i < 5; i++ {
		fmt.Println("Clássico:", i)
	}

	// Como um "while"
	j := 0
	for j < 3 {
		fmt.Println("While:", j)
		j++
	}

	// Loop infinito com break
	k := 0
	for {
		fmt.Println("Infinito, mas nem tanto...")
		if k == 2 {
			break
		}
		k++
	}

	idade := 18
	if idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
}
