package main

import (
	"fmt"
)

func issue(nr int, text string) {
	fmt.Printf("%d: %s\n", nr, text)
}

func main() {
	issue(1, "foo")
	issue(2, "bar")
	issue(123, "foobar")
}
