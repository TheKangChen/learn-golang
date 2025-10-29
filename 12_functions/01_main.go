package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// CALL BY VALUE
	num := 2.5
	nSquare := square(num)
	fmt.Printf("The square of %g is %g\n", num, nSquare)

	// CALL BY REFERENCE
	person := Person{Name: "Timothee", Age: 30}
	fmt.Printf("%s is an adult: %t\n", person.Name, IsAdult(&person))
	ShowPersonInfo(person)

	// SIDE EFFECTS: From Call By Reference
	fmt.Printf("%s's age is %d\n", person.Name, person.Age)
	fmt.Println("Happy Birthday!")
	Birthday(&person)
	fmt.Printf("Now %s's age is %d\n", person.Name, person.Age)

	ch2L, ch1L := GetLastTwoChars(person.Name)
	fmt.Printf("Last 2 characters are '%s' & '%s' \n", ch2L, ch1L)

	ch1, ch2 := GetFirstTwoChars(person.Name)
	fmt.Printf("First 2 characters are '%s' & '%s' \n", ch1, ch2)

	//BLANK IDENTIFIER: Discard values
	_ = GetNumOfChars(person.Name)

	// Passing arbitrary numbers of arguments
	PrintStrings("apple", "banana", person.Name)
	PrintStrings()
}

func square(num float64) float64 {
	// A function that returns values must end with `return` or `panic`
	return num * num
}

func ShowPersonInfo(p Person) {
	// Function without return: niladic function
	fmt.Printf("%s is %d years old\n", p.Name, p.Age)
}

func Birthday(p *Person) {
	// Passing by reference can have side effects
	p.Age += 1
}

func IsAdult(p *Person) bool {
	// Function with just the return type
	return p.Age >= 18
}

func GetLastTwoChars(s string) (string, string) {
	runes := []rune(s)
	length := len(s)
	charSecondLast := string(runes[length-2])
	charLast := string(runes[length-1])
	return charSecondLast, charLast
}

func GetFirstTwoChars(s string) (char1, char2 string) {
	// Named return variables: Defined in function signature
	runes := []rune(s)
	char1 = string(runes[0])
	char2 = string(runes[1])
	return // Named return vars can have empty return statements
}

func GetNumOfChars(s string) int {
	return len(s)
}

func PrintStrings(values ...string) {
	// Variable length arguments
	if len(values) == 0 {
		fmt.Println("No arguments provided")
	}
	for _, s := range values {
		fmt.Println(s)
	}
}
