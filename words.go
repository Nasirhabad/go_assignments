package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Worker function to count words
func countWords(text <-chan string, wordCount map[string]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for line := range text {
		words := strings.Fields(line)
		mu.Lock()
		for _, word := range words {
			wordCount[word]++
		}
		mu.Unlock()
	}
}

func main() {
	// Open the text file
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Prepare for concurrency
	var wg sync.WaitGroup
	wordCount := make(map[string]int)
	var mu sync.Mutex

	text := make(chan string, 10)

	// Start worker goroutines
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go countWords(text, wordCount, &wg, &mu)
	}

	// Read the file and send lines to the channel
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text <- scanner.Text()
	}
	close(text)

	// Wait for all goroutines to finish
	wg.Wait()

	// Create and open the CSV file
	csvFile, err := os.Create("word_frequencies.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Word", "Frequency"})

	// Write word frequencies
	for word, count := range wordCount {
		writer.Write([]string{word, fmt.Sprintf("%d", count)})
	}

	fmt.Println("Word frequencies have been written to word_frequencies.csv")
}
