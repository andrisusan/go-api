package main

import "fmt"

func main() {

	/**
	in const you have to direct initiate value
	we can compile constant even we don't use it yet
	*/
	const hello string = "Hello"
	const world = "World"
	const sample = 1 // sample unused const
	fmt.Println(hello + " " + world)

	// multiple const

	const (
		firstName = "Andri"
		lastName  = "Susanto"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
