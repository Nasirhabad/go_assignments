package main

import "fmt"

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	if n < 0 {
		return 0
	}

	a, b := 0, 1
	sum := a + b

	for i := 2; i <= n; i++ {
		next := a + b
		sum += next
		a, b = b, next
	}

	return sum
}
