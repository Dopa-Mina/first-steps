package main

import (
	"fmt"
)

func sub(a int, b int) int {
	c := a - b
	return c
}

func main() {
	difference1 := sub(10, 5)
	difference2 := sub(27, 2)
	fmt.Println(difference1, difference2)
}
