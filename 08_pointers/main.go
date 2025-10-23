package main

import "fmt"

func main() {
	// Address of operator: &
	var s string = "The name is Bond, James Bond."
	fmt.Printf("The memory address of string '%s' is: %p\n", s, &s)

	// Pointer: reference
	var pStr *string
	pStr = &s
	var pPtr *string
	pPtr = pStr
	fmt.Printf("'pStr': %p; 'pPtr': %p\n", pStr, pPtr)

	// Dereference operator *: flattens the pointer and returns the value inside the memory location
	fmt.Printf("Content of 'pStr': %s; 'pPtr': %s\n", *pStr, *pPtr)
	fmt.Println(s == *(&s))
}
