package main

import "fmt"

func main() {
	var numeros = [7] int{1,1,2,3,5,8,13}

	fmt.Println(numeros[2:5])
	fmt.Println(numeros[2:])
	fmt.Println(numeros[:5])
	fmt.Println(numeros[:])
}