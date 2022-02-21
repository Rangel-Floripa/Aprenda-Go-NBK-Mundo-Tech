package main

import (
	"fmt"
	"time"
)

func main() {

	data := time.Now()
	fmt.Println(data.Format("02/01/2006 15:04:05"))
	fmt.Println(data.Format(" 2/01/2006 15:04:05"))
	fmt.Println(data.Format(" Monday 02/01/2006 15:04:05"))
	fmt.Println(data.Format(" Mon 02/01/2006 15:04:05"))
	fmt.Println(data.Format("02/1/2006 15:04:05"))
	fmt.Println(data.Format("02/January/2006 15:04:05"))

}