package main

import (
    "fmt"
	"strconv"
)

func main() {
    salarioInicial := "1000"
	aumento := "20" 
	//salarioConvertido := 0
	//aumentoConvertido := 0

	aumentoConvertido, _ := strconv.Atoi(aumento)
	salarioConvertido, _ := strconv.Atoi(salarioInicial)
	novoSalario := ((salarioConvertido * aumentoConvertido) / 100 + salarioConvertido)

	fmt.Println(salarioConvertido)
	fmt.Println(aumentoConvertido)
    fmt.Println("O salário é", salarioInicial, "e o aumento de", aumento + "%")
	fmt.Println("O novo salário é", novoSalario)
	// para retornar o novo salário em string

	novoSalarioConvertidoString := strconv.Itoa(novoSalario)
	fmt.Println(novoSalarioConvertidoString)
}