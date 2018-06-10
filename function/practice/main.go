package main

import (
	"fmt"
)

func issueDigit(min int, max int, digit map[int]string) {
	for i := min; i <= max; i++ {
		text, ok := digit[i]
		if !ok {
			fmt.Println("digit not found")
			continue
		}
		fmt.Println(text)
	}
}

func main() {
	digit := map[int]string{1: "one", 4: "four", 9: "nine", 11: "eleven", 14: "fourteen"}
	issueDigit(1, 5, digit)
	issueDigit(8, 14, digit)
}
