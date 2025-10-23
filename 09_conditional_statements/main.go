package main

import "fmt"

func main() {
	num1 := 30

	if num1 > 0 && num1 < 20 {
		fmt.Println("0<num<20")
	} else if num1 < 30 {
		fmt.Println("20<=num<30")
	} else {
		fmt.Println("num>=30")
	}

	// You can initialize a variable (accessible only in the condition blocks) in the if statement like this:
	if num2 := 15; num2 > 10 {
		fmt.Println("num2 is greater than 10!")
		fmt.Println(num2)
	}
	// Or assign it before the conditionals
	var num3 int
	if num3 = 15; num3 > 10 {
		fmt.Println("num3 is greater than 10!")
	}

	// Switch statements: The first true condition is executed
	switch num1 {
	case 10:
		fmt.Println("'num1' is 10!")
	case 20, 25:
		fmt.Println("'num1' is 20, or 25!")
	default:
		fmt.Println("This is the default case!")
	}

	switch {
	case num3 == 10:
		fmt.Println("'num3' is 10")
	case num3 < 10:
		fmt.Println("'num3' is lesser than 10")
	case num3 > 10:
		fmt.Println("'num3' is greater than 10")
	default:
		fmt.Println("You won't be able to reach here!")
	}

	switch num4 := 100; {
	case num4 == 50:
		fmt.Println("'num4' is 50")
	case num4 > 50:
		fmt.Println("'num4' is greater than 50")
	default:
		fmt.Println("'num4' is lesser than 50")
	}

}
