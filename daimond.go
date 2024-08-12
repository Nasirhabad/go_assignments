package main

import (
	"fmt"
)

func main() {
	fmt.Println(getDiamondSeq(4))   // 16
	fmt.Println(getDiamondSeq(25))  // 625
	fmt.Println(getDiamondSeq(100)) // 10000
}

func getDiamondSeq(n int) int {
	return n * n
}
