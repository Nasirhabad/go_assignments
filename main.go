package main

import "fmt"

// Define a map to store previously computed Fibonacci numbers
var memo = make(map[int]int)

func fibonacci(number int) int {
	// Base cases
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}

	// Check if the result is already in the memo
	if result, exists := memo[number]; exists {
		return result
	}

	// Compute the Fibonacci number and store it in the memo
	memo[number] = fibonacci(number-1) + fibonacci(number-2)

	return memo[number]
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
