package main

import "fmt"

func main() {
	// Reslicing: Change the length of the slice
	// Max when reach capacity
	s := make([]int, 10, 20)

	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	fmt.Println(s)

	// Resize a slice by using the slices index from 0 to the end of the underlying array
	// [startIdx:endIdx]
	// [:] -> Entire array/slice
	s = s[:len(s)+1]
	fmt.Println(s)
}
