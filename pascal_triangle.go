package main

import (
	"fmt"
)

func generatePascalTriangle(rows int) [][]int {
	triangle := make([][]int, rows)

	for i := 0; i < rows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	var rows int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	triangle := generatePascalTriangle(rows)

	for _, row := range triangle {
		fmt.Println(row)
	}
}
