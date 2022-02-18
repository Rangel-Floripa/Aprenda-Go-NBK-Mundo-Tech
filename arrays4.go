package main

import "fmt"

func main() {
	var numeros = [7]int{1, 1, 2, 3, 5, 8, 13}
	var soma = 0
	for i := 0; i < 7; i++ {
		soma += numeros[i]
	}
	fmt.Println(soma)
}