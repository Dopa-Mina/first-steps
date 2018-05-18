package main

import (
	"fmt"
)

type person struct {
	name     string
	nachname string
}

func main() {

	p1 := person{"Max", "Mustermann"}
	p2 := person{"Moritz", "Mustermann"}
	p3 := person{"Michelle", "Mustermann"}
	p4 := person{"Manuel", "Mustermann"}

	isAdmin := make(map[person]bool)
	isAdmin[p1] = true
	isAdmin[p2] = false
	isAdmin[p3] = false
	isAdmin[p4] = false

	fmt.Println(p1, isAdmin[p1])

}
