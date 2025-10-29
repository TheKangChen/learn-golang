package main

import "fmt"

func main() {
	// defer: postpone the execution of a statement or function until the end of the enclosing function
	// Multiple defer will follow LIFO order
	defer fmt.Println("This line is last!")
	defer greet()
	defer fmt.Println("Mary had a little lamb~")
	fmt.Println("ET go home.")
}

func greet() {
	fmt.Println("Hello, universe!")
}
