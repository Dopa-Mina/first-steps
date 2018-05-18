package main

import (
	"fmt"
)

func main() {

	type person struct {
		vorname    string
		nachname   string
		strasse    string
		hausnummer int
		plz        int
		wohnort    string
		id         int
	}

	Teilnehmerliste := []person{}
	Teilnehmerliste = append(Teilnehmerliste, person{"Max", "Mustermann", "Musterstrasse", 12, 12345, "Musterstadt", 1})
	Teilnehmerliste = append(Teilnehmerliste, person{"Michaela", "Mustermann", "Musterstrasse", 1, 12345, "Musterstadt", 2})
	Teilnehmerliste = append(Teilnehmerliste, person{"Paul", "Paulsen", "Paulstrasse", 15, 12345, "Musterstadt", 3})
	fmt.Println(Teilnehmerliste)

	// range over each element in Teilnehmerliste for k = key and v = value
	for k, v := range Teilnehmerliste {
		fmt.Println(k, v)
	}
}
