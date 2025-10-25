package main

import "fmt"

func main() {
	// FUNCTION LITERALS or LAMBDA FUNCTIONS: Annonymous functions

	// Using lambdas inline
	_ = Calculate(func(x, y int) int { return x + y }, 3, 4)

	// Calling lambdas directly
	_ = func(x, y int) int { return x + y }(3, 4)

	// Assigning a lambda to a variable
	fplus := func(x, y int) int { return x + y }
	fmt.Printf("fplus is type: %T. ", fplus)
	fmt.Println("Result:", fplus(3, 4))

	// Using lambdas with defer
	defer func() { fmt.Println("This lambda is defered!") }()

	// CLOSURES: Function as return variables. Often with some kind of shared resource
	// Can be used like a function factory to create a bunch of similar functions

	// Ex: Year difference function factory
	var thisYear Year = 2025

	nextYear := YearDelta(1)
	lastYear := YearDelta(-1)
	decade := YearDelta(10)

	fmt.Println("This year is:", thisYear)
	fmt.Println("Next year is:", nextYear(thisYear))
	fmt.Println("Last year is:", lastYear(thisYear))
	fmt.Println("A decade is:", decade(thisYear))
	fmt.Println("A century is:", YearDelta(100)(thisYear)) // You can also call immediately

	// Ex: Accumulation
	acc := Accumulator()

	fmt.Print("Accumulation: ")
	fmt.Print(acc(2), " , ")
	fmt.Print(acc(10), " , ")
	fmt.Print(acc(20), " , ")
	fmt.Print(acc(100), "\n")
}

func Calculate(f func(int, int) int, x, y int) int {
	return f(x, y)
}

type Year int

func YearDelta(delta Year) func(y Year) Year {
	// The returned function will remember the delta
	return func(y Year) Year { return y + delta }
}

type Adder func(int) int

func Accumulator() Adder {
	// The returned function will have access to sum even though the outer function no longer exists
	var sum int
	return func(x int) int {
		sum += x
		return sum
	}
}
