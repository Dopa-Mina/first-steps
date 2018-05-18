package main

import (
	"fmt"
)

func main() {

	zahlen := make(map[int]string)
	zahlen[1] = "eins"
	zahlen[2] = "zwei"
	zahlen[3] = "drei"
	zahlen[4] = "vier"

	fmt.Println(zahlen)

}
