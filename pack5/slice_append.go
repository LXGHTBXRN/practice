package main

import "fmt"

func main() {
	var cars []string

	cars = append(cars, "Peugeot")
	cars = append(cars, "Mercedes")
	cars = append(cars, "Subaru")

	for i, car := range cars {
		fmt.Println(i, "-", car)
	}
}
