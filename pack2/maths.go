package main

import (
	"fmt"
	"math"
)

const (
	Pi = math.Pi
	E  = math.E
)

func main() {
	radius := 8.0
	area := Pi * radius * radius

	number := 20
	logValue := math.Log(float64(number))

	fmt.Printf("Площадь круга радиуса %.2f = %.2f\n", radius, area)
	fmt.Printf("Натуральный логарифм числа %d = %.2f\n", number, logValue)
	fmt.Printf("Константа e: %.5f\n", E)
}
