package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	city string
}

func main() {

	list := make(map[int]person)
	list[1] = person{name: "Max", age: 30, city: "Town"}
	list[2] = person{name: "Moritz", age: 29, city: "Town"}
	list[3] = person{name: "Michelle", age: 30, city: "Town"}

	fmt.Println(list[3])

}
