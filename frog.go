package main

import (
	"fmt"
)

func Frog(jumps []int) int {
	// Initialize a DP array to store the minimum cost to reach each rock
	n := len(jumps)
	dp := make([]int, n)

	// Base cases
	dp[0] = 0                        // Starting at the first rock
	dp[1] = abs(jumps[1] - jumps[0]) // Cost to jump from rock 1 to rock 2

	// Fill the DP array
	for i := 2; i < n; i++ {
		// Cost of jumping from the previous rock
		jumpOne := dp[i-1] + abs(jumps[i]-jumps[i-1])
		// Cost of jumping from the rock before the previous one
		jumpTwo := dp[i-2] + abs(jumps[i]-jumps[i-2])
		// Take the minimum of these two options
		dp[i] = min(jumpOne, jumpTwo)
	}

	// The last element of dp array will have the minimum cost to reach the last rock
	return dp[n-1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
