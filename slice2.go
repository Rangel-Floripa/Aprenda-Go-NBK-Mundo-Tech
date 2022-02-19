package main

import "fmt"

func main() {

	var nomes = [3]string{
		"Ana",
		"José",
		"Maria",
	}
	var p1[] string = nomes[0:2]
	fmt.Println(p1)
	p1[0] = "Rogério"
	fmt.Println(p1)
	//SE EU ALTERAR UM PEDAÇO DE UM ARRAY ESTAREI MODIFICANDO TAMBÉM O ARRAY ORIGINAL
	fmt.Println(nomes)
}
