package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Привет, я учу язык Go"

	words := strings.Fields(sentence)

	for _, word := range words {
		fmt.Println(word)
	}
}
