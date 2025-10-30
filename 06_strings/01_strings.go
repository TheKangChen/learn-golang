package main

import "fmt"
import "strings"
import "strconv"

func main() {
	// STRINGS: Sequence of utf-8 characters
	// 			Go strings don't use fixed width encoding, they're variable width -> more memory efficient.
	//			Since utf-8 (Unicode) is the standard, no need to decode or encode strings like in other languages

	// Interpreted Strings
	interpStr := "I say \t Hey! \t Hey!\n"

	// Raw Strings: escape characters don't get interpreted
	rawStr := `This is a raw string.\n`

	fmt.Println(interpStr)
	fmt.Println(rawStr)

	// Length
	fmt.Println("Length of 'Hello': ", len("Hello"))

	// Index: Accessing raw bytes of a string
	fmt.Println("First byte of 'Hello': ", "Hello"[0])
	fmt.Println("Last byte of 'Hello': ", "Hello"[len("Hello")-1])

	// String Concatenation
	s := "Hello, "
	s += "Benjamin " + "Button"
	fmt.Println(s)

	// STRING FUNCTIONS: https://pkg.go.dev/strings
	sentence := "A quick brown fox jumps over a box      "
	// Check substring
	fmt.Println(strings.Contains(sentence, "jump"))
	// Index of substring
	fmt.Println(strings.Index(sentence, "ox"))
	fmt.Println(strings.LastIndex(sentence, "ox"))
	// Casing
	fmt.Println(strings.ToLower(sentence))
	// Replace
	fmt.Println(strings.Replace(sentence, "fox", "tiger", -1))
	// Repeat
	fmt.Println(strings.Repeat("Hello ", 3))
	// Trimming
	fmt.Println(strings.TrimSpace(sentence))
	fmt.Println(strings.Trim(sentence, "_"))
	// Split & join
	fmt.Println(strings.Fields(sentence)) // Split on space
	fmt.Println(strings.Split(sentence, "fox"))
	fmt.Println(strings.Join(strings.Fields(sentence), "="))

	// CONVERSIONS
	fmt.Println(strconv.Itoa(10))
	strI, _ := strconv.Atoi("10")
	fmt.Println(strI)
	fmt.Println(strconv.FormatFloat(25.2, 'f', 1, 64))
	strF, _ := strconv.ParseFloat("25.2", 64)
	fmt.Println(strF)
}
