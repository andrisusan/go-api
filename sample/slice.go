package main

import "fmt"

func main() {

	// array sample
	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	var slice = days[1:4] // 1 is length, 4 is capacity
	fmt.Println(slice)
	fmt.Println("Length", len(slice))
	fmt.Println("Capacity", cap(slice))

	// slice with make
	var makeSlice = make([]string, 2, 5)
	makeSlice[0] = "Ini"
	makeSlice[1] = "Itu"

	fmt.Println(makeSlice)

	// slice sample
	var sampleSlice = []string{"test", "toast"}

	fmt.Println(sampleSlice)
}
