package main

import "fmt"

func main() {
	var nota int = 4

	switch true {
		// Posso escrever apenas switch, sem o true, ele ficará implicito
	case nota > 9:
		fmt.Println("Ótimo")
	case nota > 7:
		fmt.Println("Muito bem")
	case nota > 6:
		fmt.Println("Bom")
	default: 
		fmt.Println("Péssimo")
	}
}