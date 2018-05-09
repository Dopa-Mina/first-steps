package main

import (
	"fmt"
)

func main() {

	//create an array,

	arr := [3]int{1, 2, 3}
	fmt.Println(len(arr))
	fmt.Println(arr[1])
	fmt.Println(arr)

	//create a slice

	slice := []string{}
	slice = append(slice, "foo", "bar", "foobar")
	fmt.Println(slice)

	secondslice := []string{"what", "ever"}
	slice = append(slice, secondslice...)
	fmt.Println(slice)
}
