package main

import (
	"fmt"
	"time"
)

func main() {
	// Recursion: Functions that call itself
	var n int = 20
	start := time.Now()
	for i := 0; i <= n; i++ {
		fmt.Println(fibonacci(i))
	}
	elapsedTime := time.Since(start)
	fmt.Println("Time taken for fibonacci:", elapsedTime)

	start = time.Now()
	for i := 0; i <= n; i++ {
		fmt.Println(fibonacciMemo(i))
	}
	elapsedTime = time.Since(start)
	fmt.Println("Time taken for fibonacciMemo:", elapsedTime)
}

// Simple Recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Recursion With Memoization
var memo = make(map[int]int)

func fibonacciMemo(n int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	memo[n] = result // Cache the results

	return result
}
