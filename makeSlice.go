package main

import "fmt"

func main() {
	numeros := make([]int, 5, 7)
	fmt.Println(numeros)
	fmt.Println(cap(numeros))
	numeros[4] = 8
	fmt.Println(numeros)
}