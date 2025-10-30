package main

import (
	"fmt"
)

func main() {
	// Strings are a read-only slice of bytes (Immutable)
	s := "\u4e16\u754c"
	// When iterating through strings, you are iterating through the runes
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
	fmt.Println()

	// Converting strings to bytes -> each byte will be a unicode code point
	bs := []byte(s)
	for i, b := range bs {
		fmt.Printf("%d:%c ", i, b)
	}
	fmt.Println()

	// You can also append/concatenate strings onto bytes
	bs = append(bs, "Hello"...)
	fmt.Println(string(bs))

	// SUBSTRINGS: you can directly slice the string to get the substring if the string is perfectly one byte per character, but if not, convert it to runes first to safely slice the string based on visible characters
	sc := "世界大同"
	fmt.Println(sc[2:]) // Slicing the bytes not chars

	cc := []rune(sc)
	fmt.Println(string(cc[2:])) // Convert back to string for original characters

	// Changing characters in a string (without using the strings package)
	se := "Hello"
	ce := []byte(se)
	ce[0] = 'C'
	se2 := string(ce)
	fmt.Println(se2)
}
