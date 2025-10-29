package main

import "fmt"

func main() {
	// APPEND: returns a new slice (will always succeed unless run out of memory)
	s := make([]int, 5, 10)
	s = append(s, 20) // append to slice
	s1 := []int{100, 200}
	s = append(s, s1...) // append all items of a slice to another slice
	fmt.Println(s)

	// Removing elements from a slice using append
	s = append(s[:5], s[5+1:]...)
	fmt.Println(s)

	// COPY: Copy from source slice to destination slice
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n)

	// Append and copy are very flexible and can be used beyond just appending and copying
	// D
	// Example: Deleting an item from a slice
	// Copying sl_to so that the underlying array doesn't get modified from append
	sl_new := make([]int, len(sl_to), cap(sl_to))
	copy(sl_new, sl_to)
	sl_new = append(sl_new[:1], sl_new[1+1:]...) // returning a new slice
	fmt.Println(sl_new)

	// Deleting an item inplace
	copy(sl_to[1:], sl_to[1+1:])
	sl_to = sl_to[:len(sl_to)-1]
	fmt.Println(sl_to)
}
