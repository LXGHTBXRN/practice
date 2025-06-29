package main

import "fmt"

func incrementValue(n int) {
	n = n + 1
	fmt.Println("Внутри incrementValue", n)
}

func incrementPointer(n *int) {
	*n = *n + 1
	fmt.Println("Внутри incrementPointer", *n)
}

func main() {
	a := 5
	fmt.Println("Исходное значение: ", a)

	incrementValue(a)
	fmt.Println("После incrementValue: ", a)

	incrementPointer(&a)
	fmt.Println("После incrementPointer: ", a)
}
