package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	// `sort` package provides methods to sort and search (binary) user defined collections, including slices/arrays
	arrI := []int{4, 2, 3, 1, 0}
	arrF := []float64{2.5, 7.5, 1.6, 10.0, 2.8}
	arrS := []string{"banana", "kiwi", "apple", "lemon", "dragon fruit"}

	sort.Ints(arrI)     // Calles slices.Sort() underneath the hood in go 1.22
	sort.Float64s(arrF) // Calles slices.Sort() underneath the hood in go 1.22
	slices.Sort(arrS)

	fmt.Println(arrI)
	fmt.Println(arrF)
	fmt.Println(arrS)

	// Binary search
	idx := sort.SearchInts(arrI, 3)
	fmt.Println(idx)
	idx = sort.SearchInts(arrI, 10) // returns index to insert if doesn't exist
	fmt.Println(idx)

	idx = sort.SearchFloat64s(arrF, 10)
	fmt.Println(idx)

	idx = sort.SearchStrings(arrS, "apple")
	fmt.Println(idx)
}
