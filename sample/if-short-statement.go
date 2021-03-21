package main

import "fmt"

func main() {

	var airport = "Soekarno Hatta"

	if length := len(airport); length > 2 {
		fmt.Println("length more than 2 characters")
	} else {
		fmt.Println("length less than 2 characters")
	}

}
