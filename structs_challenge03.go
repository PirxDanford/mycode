package main

import "fmt"

type Astro struct {
	name string
	age int
	lastmission string
	isneeded bool
}

type nasaMission struct {
	people []Astro
	number int
	message string
}

func main() {
	entry1 := Astro{"Pirx Danford",48,"Go Training Day 2",true}
	entry2 := Astro{"Helen van Scharl",28,"Dreams among Stars",false}
	entry3 := Astro{"Matthias Maurer",52,"Cosmic Kiss",false}
	AstroSlice := []Astro{entry1,entry2,entry3}
	mission := nasaMission{AstroSlice,3,"success"}
	fmt.Println(mission)
	fmt.Printf("%+v",mission)
}
