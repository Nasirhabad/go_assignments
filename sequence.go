package main

import "fmt"

func main() {
	fmt.Println(getSequence(4))   // 10
	fmt.Println(getSequence(15))  // 120
	fmt.Println(getSequence(100)) // 5050
}

func getSequence(n int) int {
	// Calculate the sum of the first n positive integers
	sum := n * (n + 1) / 2
	// Multiply by n to get the final result
	return n * sum
}
