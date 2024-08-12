package main

import "fmt"

func main() {
	primeFactors(20)   // Output: 2, 2, 5
	primeFactors(75)   // Output: 3, 5, 5
	primeFactors(1024) // Output: 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {
	for n%2 == 0 {
		fmt.Print(2, " ")
		n /= 2
	}
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			fmt.Print(i, " ")
			n /= i
		}
	}
	if n > 2 {
		fmt.Print(n, " ")
	}
	fmt.Println()
}
