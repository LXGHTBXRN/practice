package main

import "fmt"

func main() {
	var a, b float64

	fmt.Println("Введите два числа:")
	fmt.Scan(&a, &b)

	fmt.Printf("Сложение: %f + %f = %f\n", a, b, a+b)
	fmt.Printf("Вычитание: %f - %f = %f\n", a, b, a-b)
	fmt.Printf("Умножение: %f * %f = %f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("Деление: %f / %f = %f\n", a, b, a/b)
		fmt.Printf("Остаток: %f %% %f = %f\n", a, b, float64(int(a)%int(b)))
	} else {
		fmt.Println("Деление на ноль невозможно")
	}
}
