package main

import "fmt"

func main() {
	// Higher-Order Functions: Functions that take or return other functions (function as values)
	fmt.Println(Calculate(Add, 2, 3))
	fmt.Println(Calculate(Mult, 2, 3))
}

func Calculate(f func(int, int) int, x, y int) int {
	// Higher-order math function that performs calculations
	return f(x, y)
}

func Add(x, y int) int {
	return x + y
}

func Mult(x, y int) int {
	return x * y
}
