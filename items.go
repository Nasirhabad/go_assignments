package main

import (
	"fmt"
)

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {
	var result []int
	findCombination(items, target, 0, []int{}, &result)
	return result
}

func findCombination(items []int, target, start int, currentCombination []int, result *[]int) {
	if target == 0 {
		*result = append([]int{}, currentCombination...)
		return
	}

	for i := start; i < len(items); i++ {
		if items[i] <= target {
			findCombination(items, target-items[i], i+1, append(currentCombination, items[i]), result)
			if len(*result) > 0 {
				return
			}
		}
	}
}
