package main

import "fmt"

func main() {

	// declare variable without var keyword but must initial variable directly
	a := 20
	b := 10

	fmt.Println("Addition ", a+b)
	fmt.Println("Minus ", a-b)
	fmt.Println("Multiply ", a*b)
	fmt.Println("Division ", a/b)
	fmt.Println("Modulus ", a%b)

	// augmented assignment
	a += 10
	fmt.Println(a)

	// unary operator
	a++
	fmt.Println(a)
}
