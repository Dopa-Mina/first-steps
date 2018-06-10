package main

import (
	"fmt"
)

func main() {

	for x := 1; x <= 10; x++ {
		for y := 1; y <= 10; y++ {
			fmt.Printf("%02d:%02d\n", x, y)
			if y == x {
				break
			}
		}
		fmt.Println("-----")
	}
}
