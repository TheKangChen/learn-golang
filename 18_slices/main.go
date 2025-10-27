package main

import "fmt"

func main() {
	// Slices are reference to a continuous segment of an array and unlike an array the length of a slice can change (min 0, max length of underlying array)

	// Declaring a slice
	var a1 = [...]int{1, 2, 3, 4, 5}
	var s1 []int = a1[1:3]

	var a2 [3]string
	var s2 = a2[:] // Slice can be defined before an array is populated

	// Slices can also be initialized like arrays
	s3 := []int{10, 20, 30}
	var _ = s3[:] // Slice from another slice

	// Slices are small 3 property data structure: address to underlying array, capacity, and length
	fmt.Println("s1:", s1)
	fmt.Println("capacity of s1:", cap(s1)) // capacity of an slice is from the start of the slice to the end of the underlying array
	fmt.Println("length of s1:", len(s1))

	s2[1] = "banana" // modifies underlying array since it's an reference
	fmt.Println("a2:", a2)

	// Creating a slice with preallocated memory using `make` (it'll also create the underlying array)
	var _ []int = make([]int, 5)

	// `make` initializes and returns an object
	s4 := make([]int, 5, 10)
	// `new` allocates memory, zeros it, then return the memory locaiton (like calloc in C)
	s5 := new([10]int)[0:5]
	fmt.Println(s4)
	fmt.Println(s5)

	// This is a nil slice. It's not referencing anything yet.
	var s6 []int
	fmt.Println(s6, len(s6), cap(s6)) // output: [] 0 0

	// Passing slices instead of array as function args
	sum := Sum(a1[:])
	fmt.Println("The sum of a1:", sum)
}

func Sum(s []int) (sum int) {
	// For functions that takes an array, it's often better for the parameter to be a slice
	for _, v := range s {
		sum += v
	}
	return
}
