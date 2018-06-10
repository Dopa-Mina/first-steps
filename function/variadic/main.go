package main

import (
	"fmt"
)

func add(i ...int) int {
	var sum int
	for _, v := range i {
		sum = sum + v
	}
	return sum

}

func main() {
	fmt.Println(add(1, 2, 3))
}
