package main

import "fmt"

func main() {
	// CONSTANTS
	const pi = 3.1415926         // implicitly typed (untyped)
	const greet string = "G'day" // explicitly typed
	const calc1 = 3 / 2          // Constants may be declared as a calculation, but all vars need to available at compile time
	// Untyped constants will not overflow
	const largeNumber = 0.693147180559945309417232121458176568075500134360255254120680009
	// Multiple assignment
	const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday int = 1, 2, 3, 4, 5, 6
	const Beef, Two, C = "meat", 2, "veg"

	//Enums
	type Month int

	const (
		_ Month = iota
		Jan
		Feb
		Mar
		Apr
		May
		Jun
		Jul
		Aug
		Sep
		Oct
		Nov
		Dec
	)

	const (
		System = "system"
		Human  = "human"
		AI     = "assistant"
	)

	fmt.Println(Jan, Feb)
	fmt.Println(System, Human)

	// VARIABLES
	var decision bool // Variable declaration
	fmt.Println(decision)

	var someNumber float32 = 25.12 // Initializing a variable: declare and assign value to it
	var isLearningGo = true        // Automatic type inference
	fmt.Println(someNumber)
	fmt.Println(isLearningGo)

	// ASSIGNMENT OPERATOR := or INITIALIZING DECLARATION
	name := "Kang"
	fmt.Println(name)
}
