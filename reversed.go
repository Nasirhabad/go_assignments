package main

import (
	"fmt"
	"time"
)

func reverseWord(word string) {
	for i := 1; i <= len(word); i++ {
		go func(subWord string) {
			fmt.Println(subWord)
		}(reverse(word[:i]))
		time.Sleep(3 * time.Second)
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reverseWord("phone")
}
