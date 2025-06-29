package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"

	lenght := len([]rune(str))
	fmt.Println("Количество символов:", lenght)

	substr := "Go"
	contains := strings.Contains(str, substr)
	fmt.Println("Строка содержит - ", substr, contains)

	upper := strings.ToUpper(str)
	fmt.Println("Верхний регистр: ", upper)

	lower := strings.ToLower(str)
	fmt.Println("Нижний регистр: ", lower)
}
