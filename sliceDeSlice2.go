package main

import (
	"fmt"
	"strings"
)

func main() {
	tabuleiro := [][]string{
		[]string{"x", "0", "x"},
		[]string{"0", "x", "0"},
		[]string{"0", "0", "x"},
	}
	
	fmt.Println(tabuleiro[0])
	fmt.Println(tabuleiro[1])
	fmt.Println(tabuleiro[2])
	fmt.Println(tabuleiro)
	fmt.Println(tabuleiro[0][0])

	for i := 0; i < len(tabuleiro); i++ {
		fmt.Printf("%s\n", strings.Join(tabuleiro[i], " "))
	}

}
