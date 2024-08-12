package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}

	// Create big integers to handle large factorials
	twoNFactorial := new(big.Int).MulRange(1, int64(2*n))
	nFactorial := new(big.Int).MulRange(1, int64(n))
	nPlusOneFactorial := new(big.Int).MulRange(1, int64(n+1))

	// Catalan number formula: (2n)! / ((n+1)! * n!)
	catalanNumber := new(big.Int).Div(twoNFactorial, nPlusOneFactorial)
	catalanNumber.Div(catalanNumber, nFactorial)

	return catalanNumber
}
