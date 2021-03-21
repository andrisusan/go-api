package main

import "fmt"

func main() {

	var value32 int32 = 128
	var value64 int64 = int64(value32)

	// when convert with over limit, will be begin from beginning again
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8) // this result will -128 cause over limit int8

	var name = "Andri"
	var a byte = name[0]           // result will be in byte
	var aString string = string(a) // convert byte to string to see what char is

	fmt.Println(name)
	fmt.Println(aString)
}
