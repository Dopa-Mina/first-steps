package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")

		case i%3 == 0:
			fmt.Println("Fizz")

		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}

}

/*
	Erstelle ein Programm, welches die Zahlen von 1 bis 100
	ausgibt. Wenn die Zahl jedoch durch 3 teilbar ist, dann
	gib Fizz aus. Wenn die Zahl durch 5 teilbar ist, dann gib
	Buzz aus. Für Zahlen, welche durch 3 und durch 5 teilbar
	sind (z.B. 15) gib FizzBuzz aus.

*/
