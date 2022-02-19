package main

import "fmt"

func main() {
	var animes [3]string
	animes = [3]string{"Dragon  Ball", "Sailor Moon", "Yuyu Hakusho"}
	fmt.Println(animes)
	var p1 = animes[0:2]
	fmt.Println(p1)
	fmt.Println(cap(p1))
	fmt.Println(len(p1))
	fmt.Printf("cap=%d len=%d\n", cap(p1), len(p1))

	var p2 =animes[1:]
	fmt.Println(p2)
	fmt.Printf("cap=%d len=%d\n", cap(p2), len(p2))
}