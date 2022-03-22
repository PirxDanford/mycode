package main

import "fmt"

type Astro struct {
	name string
	age int
	lastmission string
	isneeded bool
}

func main() {
	entry1 := Astro{"Pirx Danford",48,"Go Training Day 2",true}
	entry2 := Astro{"Helen van Scharl",28,"Dreams among Stars",false}
	entry3 := Astro{"Matthias Maurer",52,"Cosmic Kiss",false}
	fmt.Println(entry1)
	fmt.Println(entry2)
	fmt.Println(entry3)
}
