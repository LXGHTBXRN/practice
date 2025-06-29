package main

import "fmt"

func main() {
	var year int
	fmt.Println("Введите год:")
	fmt.Scan(&year)

	isLeap := (year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeap {
		fmt.Printf("Год %d является весокосным\n", year)
	} else {
		fmt.Printf("Год %d не является весокосным\n", year)
	}
}
