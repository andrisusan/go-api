package main

import "fmt"

func main() {
	var a = "Hello"
	var b = "Hello go"

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(len(a))
	fmt.Println(len(b))

	fmt.Println("index 0 in string a in byte", a[0])
	fmt.Println("index 1 in string b in byte", b[1])

	a = "Hello bro"
	fmt.Println(a)

}
