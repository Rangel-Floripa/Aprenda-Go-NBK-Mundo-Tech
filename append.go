package main

import "fmt"

func main() {
	var nomes [] string

	fmt.Println(nomes)
	var nomes2 = append(nomes, "João")
	nomes = append(nomes2, "Maria")
	fmt.Println(nomes)
	fmt.Println(nomes2)
	fmt.Printf("len= %d ; cap= %d\n", len(nomes2), cap(nomes2))
	fmt.Printf("len= %d ; cap= %d\n", len(nomes), cap(nomes))

}