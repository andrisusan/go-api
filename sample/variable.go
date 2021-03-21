package main

import "fmt"

func main() {
	// declare without data type
	var a = "Hello"

	// declare variable with data type
	var b string
	b = "Hello go"

	// declare variable without var keyword
	c := "Hello go world"

	// declare multiple variable
	var (
		hello = "Hello"
		world = "World"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(hello)
	fmt.Println(world)

}
