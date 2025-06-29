package main

import "fmt"

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x, y := 1, 2
	fmt.Printf("До swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("После swap: x = %d, y = %d\n", x, y)
}
