package main

import "fmt"

// User defined data types
type IZ int

func main() {
	// PRIMATIVE TYPES
	// Integers: int, uint (8, 16, 32, 64 bits)
	var negTen int = -10
	var _ uint = 100
	var _ byte = 'A' // uint8
	var _ rune       // int32 A unicode code point (utf-8 characters)

	// Floating-Point Numbers
	var number float32 = 20.1
	var _ float64 = 34.12345678

	// Complex Numbers
	var complexNumber complex64 = 1 + 4i
	var _ complex128 = complex(5.0, 1.5) // (real, imaginary)

	// Strings
	var message string = "I love golang!"

	// BOOLEANS
	var decision bool = true

	fmt.Println(number)
	fmt.Println(IZ(number)) // Type casting from float to int
	fmt.Println(complexNumber)
	fmt.Println(message)
	fmt.Println(decision)

	// TYPE CONVERSION
	fmt.Println(number)
	fmt.Println(float64(number))

	// FORMAT SPECIFIERS
	fmt.Printf("Value of negTen: %03d\n", negTen)             // int
	fmt.Printf("Value of number: %7.5g\n", number)            // float - e: scientific notation, f: float 7 digits (like C)
	fmt.Printf("Value of complexNumber: %v\n", complexNumber) // complex
	fmt.Printf("Value of message: %s\n", message)             // string
	fmt.Printf("Value of decision: %t\n", decision)           // bool
}
