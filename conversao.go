package main

import (
	"fmt"
)

func main() {
	var a int = 46
	var b float64 = 6.4
	var c int = int(b)
	var d float64 = float64(a)
	
	fmt.Printf("Var a tipo:%T\n Var a valor:%v\n", a, a)
	fmt.Printf("Var b tipo:%T\n Var b valor:%v\n", b, b)
	fmt.Printf("Var c tipo:%T\n Var c valor:%v\n", c, c)
	fmt.Printf("Var d tipo:%T\n Var d valor:%v\n", d, d)
}