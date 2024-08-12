package main

import (
	"fmt"
)

// ConvertToRoman converts an integer to a Roman numeral.
func ConvertToRoman(num int) string {
	// Define Roman numeral symbols and their values.
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var roman string
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			roman += symbols[i]
		}
	}
	return roman
}

func main() {
	// Test cases
	testNumbers := []int{4, 9, 23, 2021, 1646}

	for _, num := range testNumbers {
		fmt.Printf("Input: %d\nOutput: %s\n\n", num, ConvertToRoman(num))
	}
}
