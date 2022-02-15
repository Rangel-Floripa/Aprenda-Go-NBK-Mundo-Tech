package main

import "fmt"

func main() {

	var soma int = 2
	for soma < 600 {
		soma += soma
		fmt.Println("variável soma dentro do loop", soma)
	}

	fmt.Println("valor final da variável soma", soma)
}