package main

import (
	"fmt"
	"sort"
)

func findElements(slice []string, target string) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func sortSlice(slice []string) {
	sort.Strings(slice)
}

func filterSlice(slice []string, minLen int) []string {
	var filtered []string
	for _, v := range slice {
		if len(v) > minLen {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	cars := []string{"BMW", "Mercedes", "McLaren", "Audi"}

	index, found := findElements(cars, "Mercedes")
	if found {
		fmt.Printf("Найден элемент 'Mercedes' на позиции %d\n", index)
	} else {
		fmt.Println("Элемент 'Mercedes' не найден")
	}

	sortSlice(cars)
	fmt.Println("Отсортированный срез", cars)

	filtered := filterSlice(cars, 4)
	fmt.Println("Отсфильтрованный срез (длина > 4):", filtered)
}
