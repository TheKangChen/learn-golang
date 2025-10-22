package main

import "fmt"

var number int = 10 // Global scope

func main() {
	var decision bool = true // Local (function) scope

	fmt.Printf("Original value of number: %d\n", number) // Format string
	number = 15 // reassignment

	fmt.Printf("New value of number: %d\n", number)
	fmt.Printf("Decision: %t\n", decision)
}
