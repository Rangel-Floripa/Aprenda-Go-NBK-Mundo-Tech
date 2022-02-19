package main

import "fmt"

func main() {

	var nomes = []string{
		"Ana",
		"José",
		"Maria",
	}
	fmt.Println(cap(nomes))
	fmt.Println("len =", len(nomes))
	fmt.Printf("%T\n",nomes)
	fmt.Printf("len=%d\n" , len(nomes))
	fmt.Printf("len=%d, cap=%d\n", len(nomes), cap(nomes))
}
