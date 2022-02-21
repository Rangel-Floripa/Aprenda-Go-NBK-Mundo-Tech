package main

import "fmt"

func main() {

	var numeros = []int{10, 20, 30, 40}

	for i := 0; i < len(numeros); i++ {
		fmt.Println(i)
		fmt.Println(numeros[i])
	}

	for i, v := range numeros{
		fmt.Println(i,v)
	}
}