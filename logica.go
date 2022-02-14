package main

import "fmt"

func main() {
	var x int = 9
	fmt.Println("x > 3 && x < 7", x > 3 && x < 7)
	fmt.Println("x < 3 || x > 7", x < 3 || x > 7)
	fmt.Println(!false)
}