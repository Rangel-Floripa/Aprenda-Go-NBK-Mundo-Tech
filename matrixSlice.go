package main

import "fmt"


func main() {

	tabuleiro := [][]string {
	{"x", "0", "x"},
	{"0", "x", "0"},
	{"0", "0", "x"},
	}
	
	for i:= 0 ; i < len(tabuleiro); i++ {
		for j:= 0; j < len(tabuleiro[i]); j++ {
			fmt.Println(tabuleiro[i][j])
		}
	}
}