package main

import (
	"fmt"
	"sync"
)

// isComposite checks if a number is composite
func isComposite(n int) bool {
	if n <= 1 {
		return false
	}
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count > 2
}

// generateComposites generates composite numbers from 1 to 100
func generateComposites(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; i++ {
		if isComposite(i) {
			out <- i
		}
	}
	close(out)
}

// squareNumbers calculates the square of each composite number
func squareNumbers(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		out <- num * num
	}
	close(out)
}

// checkEvenOdd checks if a number is even or odd and prints "Pong" or "Ping"
func checkEvenOdd(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		if num%2 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println("Ping")
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Channels for communication between goroutines
	compositeChan := make(chan int)
	squareChan := make(chan int)

	// Start goroutines
	wg.Add(1)
	go generateComposites(compositeChan, &wg)

	wg.Add(1)
	go squareNumbers(compositeChan, squareChan, &wg)

	wg.Add(1)
	go checkEvenOdd(squareChan, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
