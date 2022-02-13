package main

import "fmt"

func calcular(a int) (int, int) {
	var quadrado int = a * a
	var cubo int = a * a * a

	return quadrado, cubo
}

func main() {
	fmt.Println(calcular(4))
}