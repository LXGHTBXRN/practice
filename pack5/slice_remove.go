package main

import "fmt"

func removeIndex(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func main() {
	cars := []string{"BMW", "Mercedes", "McLaren"}
	fmt.Println("Исходный срез:", cars)

	cars = removeIndex(cars, 1)
	fmt.Println("После удаления элемента с индексом 1:", cars)
}
