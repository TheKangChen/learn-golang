package main

import "fmt"

func main() {
	// COUNTER-CONTROLLED ITERATION:
	// Syntax:
	// for initialization; condition; modification { }
	str := "Sun and Moon"
	for i := 0; i < len(str); i++ {
		fmt.Printf("Char on pos %-2d is: %c\n", i, str[i])
	}

	// Initialize multiple counters
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %d; j: %d \n", i, j)
	}

	// CONDITION-CONTROLLED ITERATION:
	// Syntax:
	// for condition { }
	var count int = 0
	for count < 5 {
		fmt.Printf("This is the iteration: %d\n", count+1)
		count++
	}

	// INFINITE LOOP:
	// Syntax:
	// for { }
	// for true { }
	// for i :=0; ; i++ { }
	count = 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println(count)
		count++
	}

	// LOOP OVER A RANGE:
	// Syntax:
	// for ix, val := range collection { }
	for ix, char := range str {
		fmt.Printf("Char on pos %-2d is: %c\n", ix, char)
	}
}
