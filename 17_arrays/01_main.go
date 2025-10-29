package main

import "fmt"

func main() {
	// Arrays are fixed sized collection of homogeneous items
	var arr1 [5]int // Must specify size!

	// Array Literals
	arr2 := [3]int{1, 2, 3}
	var _ = [10]int{1, 2, 3} // First 3 items
	var arr3 = [...]int{10, 20, 30, 40} // Compiler counts for you
	var _ = [5]string{3: "Chris", 4: "Ron"}

	// Arrays are value type, if you want it to be like C where the variable is a pointer to the first element of an arry, you need to do it like this:
	arr4 := new([5]int)
	arr5 := &arr3

	fmt.Println(arr4)
	fmt.Println(arr5)

	// Assigning values
	arr1[len(arr1)-1] = 10

	// Looping over an array
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Index %d: %d\n", i, arr1[i])
	}

	fmt.Println(arr2)
	for ix, v := range arr2 {
		fmt.Printf("Index %d: %d\n", ix, v)
	}
}
