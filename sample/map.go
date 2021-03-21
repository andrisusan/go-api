package main

import (
	"fmt"
)

func main() {

	var device map[string]string = map[string]string{
		"code": "A",
		"name": "phone",
	}

	// add key and value in map or change value
	device["id"] = "123"

	fmt.Println(device)
	fmt.Println(device["code"])
	fmt.Println(device["name"])
	fmt.Println(device["id"])

	// another way to make a map using make keyword
	var notebook map[string]string = make(map[string]string)

	notebook["name"] = "Macbook pro"
	notebook["size"] = "15 inch"
	notebook["color"] = "Grey"
	fmt.Println(notebook)

	// delete in map
	delete(notebook, "color")
	fmt.Println(notebook)
}
