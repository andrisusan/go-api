package main

import "fmt"

func main() {
	var fruit [2]string

	// initial array value
	fruit[0] = "Mango"
	fruit[1] = "Banana"

	//get array data
	a := fruit[0]
	fmt.Println(a)
	fmt.Println(fruit[1])

	// declaration array with initial values
	var animals = [2]string{
		"Giraffe",
		"Panda",
	}
	fmt.Println(animals)

	//get array size
	fmt.Println("Animal array size", len(animals))

	//array with dynamic size
	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	fmt.Println(days)
}
