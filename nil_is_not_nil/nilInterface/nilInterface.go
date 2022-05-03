package main

import (
	"fmt"
	"log"
)

func main() {
	//START OMIT
	var a *int
	var b interface{}
	var c interface{}
	b = a

	if a == nil {
		fmt.Println("a is nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	}
	if c == nil {
		fmt.Println("c is nil")
	}
	//END OMIT

	log.Println(a, b)
}
