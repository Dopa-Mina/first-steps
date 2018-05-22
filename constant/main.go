package main

import (
	"fmt"
)

const (
	S = iota
	M
	L
	XL
)

const (
	blue = iota
	green
	yellow
	red
)

type shirt struct {
	size  int
	color int
}

func main() {

	shirt1 := shirt{size: S, color: blue}
	shirt2 := shirt{size: M, color: red}

	if shirt1.size > shirt2.size {
		fmt.Println("The blue shirt is bigger")
	} else {
		fmt.Println("The red shirt is bigger")
	}
}
