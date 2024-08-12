package main

import (
	"fmt"
	"math"
)

func generatePrimes(limit int, primesChan chan<- int) {
	defer close(primesChan)
	for num := 2; num <= limit; num++ {
		if isPrime(num) {
			primesChan <- num
		}
	}
}
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func calculateSquare(primesChan <-chan int, squaresChan chan<- int) {
	defer close(squaresChan)
	for prime := range primesChan {
		squaresChan <- prime * prime
	}
}
func main() {
	primesChan := make(chan int)
	squaresChan := make(chan int)
	go generatePrimes(100, primesChan)
	go calculateSquare(primesChan, squaresChan)
	for square := range squaresChan {
		fmt.Println(square)
	}
}
