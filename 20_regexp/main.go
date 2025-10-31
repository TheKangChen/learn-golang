package main

import (
	"fmt"
	"regexp"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	// Match(): find match in byte string
	ok, err := regexp.Match(pat, []byte(searchIn))
	if err == nil {
		fmt.Println(ok)
	}

	// Compile(): parses a regular expression and returns a regexp object if successful
	// MustCompile(): panic if failed
	re, err := regexp.Compile(pat)

	// Find() & FindAll(): return matching slices of bytes
	// FindString() & FindAllString(): strings
	if err == nil {
		matchedBytes := re.FindAll([]byte(searchIn), -1)
		for _, b := range matchedBytes {
			fmt.Printf("%s ", b)
		}
		fmt.Println()

		matchedStrings := re.FindAllString(searchIn, -1)
		fmt.Println(matchedStrings)
	}

	// FindStringSubmatch():
	// If there are capture groups in the regex pattern, use FindSubmatch() to access them
	// Index: [0] full match; [1:] groups
	re = regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))

	// ReplaceAllString() & ReplaceAllStringFunc(): Replacing matched strings with string or formating function that returns a string
	re = regexp.MustCompile(pat)
	str := re.ReplaceAllString(searchIn, "##.##")
	fmt.Println(str)
}
