package main

import "fmt"

func main() {
	var nome string = "João"
	switch nome {
	case "Ana":
		fmt.Println("É a Ana")
	case "João":
		fmt.Println("É o João")
	default:
		fmt.Println("Não conheço")
	}

}