package main

import "fmt"

func main(){
	var soma int = 0

	for i := 1; i <= 10; i++ {
		soma += i
		fmt.Println("variável soma dentro do loop for = ", soma)
		fmt.Println("variável i dentro do loop for = ", i)
	}
		fmt.Println("Variável soma ==", soma)
}