package main

import "fmt"

func SimpleEquations(a, b, c int) {
	found := false
	for x := 1; x <= 100; x++ {
		for y := x + 1; y <= 100; y++ { // y > x ensures y is different from x
			for z := y + 1; z <= 100; z++ { // z > y ensures z is different from y and x
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)    // no solution
	SimpleEquations(6, 6, 14)   // 1 2 3
	SimpleEquations(12, 36, 50) // 2 4 6 (example case)
}
