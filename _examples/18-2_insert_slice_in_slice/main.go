package main

import "fmt"

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := insertSlice(s, in, 0)
	fmt.Println(res)
	res = insertSlice(s, in, 3)
	fmt.Println(res)
}

func insertSlice(slice, insertion []string, index int) []string {
	res := make([]string, len(slice)+len(insertion))
	copy(res, slice[:index])
	copy(res[index:], insertion)
	copy(res[index+len(insertion):], slice[index:])
	return res
}
