package main

import "fmt"

func main() {
	numeros := []int{1, 1, 2, 3, 5, 8}

	for indice :=0; indice < len(numeros); indice++ {
		fmt.Println(numeros[indice])
	}
	
}