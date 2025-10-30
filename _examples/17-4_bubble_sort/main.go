package main

import "fmt"

func main() {
	s := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", s)
	bubbleSort(s)
	fmt.Println("after sort:  ", s)
}

func bubbleSort(s []int) []int {
	// Add flag to track each iteration, if an iteration doesn't have any swaps, the slice is sorted -> early exit
	swapped := true
	for pass := 0; pass < len(s) && swapped; pass++ {
		swapped = false
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
	}
	return s
}
