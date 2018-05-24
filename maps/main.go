package main

import (
	"fmt"
)

func main() {

	digit := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
	}
	for i := 1; i <= len(digit); i++ {
		fmt.Println(digit[i])
	}
}
