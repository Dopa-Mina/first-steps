package main

import (
	"fmt"
	"log"
)

func main() {

	var name = "Yasemina"
	var alter = 29

	if alter == 29 {
		//log.Fatal("you shall not pass")
		log.Printf("user ist ganz schön alt nämlich: %d \n", alter)
	}

	fmt.Printf("Mein Name ist %s \n", name)
	fmt.Printf("und ich bin %v Jahre alt \n", alter)

}
