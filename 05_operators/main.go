package main

import "fmt"

func main() {
	// ARITHMETIC OPERATORS
	fmt.Println(2 + 2)
	fmt.Println(2 - 1)
	fmt.Println(2 * 3)
	fmt.Println(10 / 7)   // Divide using int result in int
	fmt.Println(10.0 / 7) // 10.0 is float so untyped constant 7 gets coerced into float
	fmt.Println(10 % 7)

	// Arithmetic Assignment
	x := 10
	x += 5
	fmt.Println(x)

	// Increment & Decrement
	y := 20.0
	y--
	fmt.Println(y)

	// LOGICAL OPERATORS
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)
	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 >= 10)
	fmt.Println(10 <= 10)

	// The values that can be compared NEEDS to of the same type!!
	fmt.Println(y > float64(x))

	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(true && !true)

	// BITWISE OPERATORS
	// NOTE: Come back later when I need it. Can't remember this now
}
