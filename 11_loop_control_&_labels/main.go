package main

import "fmt"

func main() {
	// labels: any line of code that starts with `for`, `switch`, `select` can be preceeded with a label
	str := "abc"
OuterLoop:
	for i := 0; i <= 5; i++ {
		for _, ch := range str {
			if i == 2 {
				continue // Continue from the innermost loop
			}
			if i == 3 {
				break // Break from the innermost loop
			}
			if i == 5 {
				continue OuterLoop // Break from label instead of innermost loop
			}
			fmt.Println(i)
			fmt.Printf("  %c\n", ch)
		}
	}
	goto_example()
}

func goto_example() {
	// goto can lead to poor code and bad design. Thus discouraged.
	// Use the goto LABEL only with LABEL defined after the goto line!
	i := 0
HERE:
	fmt.Printf("%d", i)
	i++
	if i == 5 {
		return
	}
	goto HERE // Can only goto labels in current scope
}
