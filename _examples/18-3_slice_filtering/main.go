package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var res []int

	res = Filter(s, even)
	fmt.Println(res)

	res = Filter(s, odd)
	fmt.Println(res)
}

func Filter(s []int, fn func(int) bool) []int {
	var res []int
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func odd(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}
