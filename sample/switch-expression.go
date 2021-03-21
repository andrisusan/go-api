package main

import (
	"fmt"
	"strings"
)

func main() {

	size := 10

	switch size {
	case 1:
		fmt.Println("not 1")
	case 2:
		fmt.Println("not 2")
	default:
		fmt.Println("size is :", size)
	}

	// another way switch expression
	switch name := "Andri"; strings.Contains(name, "a") { // sensitive case for contains function
	case true:
		fmt.Println("name contains character a lower case")
	case false:
		fmt.Println("name doesn't contains character a lower case")
	}

	// switch case like if statement
	switch {
	case size > 10:
		fmt.Println("size more than 10")
	case size < 10:
		fmt.Println("size less than 10")
	default:
		fmt.Println("size is 10")

	}

}
