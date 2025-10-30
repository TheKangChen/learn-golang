package main

import "fmt"

func main() {
	str := "Golang"
	str2 := "世界大同"

	fmt.Printf("The reversed string for -%s- is -%s-\n", str, reverse(str))
	fmt.Printf("The reversed string for -%s- is -%s-\n", str2, reverse(str2))
}

func reverse(str string) string {
	runes := []rune(str)
	n, h := len(runes), len(runes)/2
	for i := range h {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	str = string(runes)
	return str
}
