package main

import (
	"fmt"
)

func sub(a int, b int) int {
	c := a - b
	return c

}
func add(d int, f int) int {
	g := d + f
	return g

}
func div(h int, i int) int {
	j := h / i
	return j

}
func mul(k int, l int) int {
	j := k * l
	return j

}

func main() {
	diff1 := sub(10, 5)
	diff2 := sub(27, 2)
	diff3 := add(28, 4)
	diff4 := add(29, 3)
	diff5 := div(300, 15)
	diff6 := div(210, 10)
	diff7 := mul(23, 13)
	diff8 := mul(7, 8)

	fmt.Println(diff1, diff2, diff3, diff4, diff5, diff6, diff7, diff8)
}
