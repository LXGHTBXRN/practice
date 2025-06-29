package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	freq := make(map[string]int)

	words := strings.Fields(text)
	for _, word := range words {
		word = strings.ToLower(word)
		freq[word]++
	}
	return freq
}

func main() {
	text := "Я учусь на 2 курсе. Сейчас я прохожу учебную практику в своем университете."

	freqMap := wordFrequency(text)
	fmt.Println("Частоты слов:")

	for word, count := range freqMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}
