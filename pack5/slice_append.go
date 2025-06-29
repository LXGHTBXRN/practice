package main

import "fmt"

func main() {
	var cars []string

	cars = append(cars, "BMW")
	cars = append(cars, "Mercedes")
	cars = append(cars, "McLaren")

	for i, car := range cars {
		fmt.Println(i, "-", car)
	}
}
