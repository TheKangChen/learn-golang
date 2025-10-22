package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// Pseudo random: automatically seeded at start up
	x := rand.Int()              // random int
	y := rand.IntN(10)           // random int between [0, 10]
	z := rand.Float64()          // random float between[0.0, 1.0]
	w := rand.N(5 * time.Second) // generic function for any integer type
	number := getRandInt(10, 20)

	fmt.Printf("x: %d\n", x)
	fmt.Printf("y: %d\n", y)
	fmt.Printf("z: %g\n", z)
	fmt.Printf("w: %d\n", w)
	fmt.Printf("number: %d\n", number)
}

func getRandInt(min, max int) int {
	// Include min and max
	return rand.IntN(max-min+1) + min
}
