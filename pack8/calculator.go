package main

import "fmt"

func main() {
	var operator string
	fmt.Println("Введите оператор: +, -, *. /")
	fmt.Scan(&operator)

	var num1, num2 float64
	fmt.Println("Введите два числа:")
	fmt.Scan(&num1, &num2)

	switch operator {
	case "+":
		fmt.Println("Результат: ", num1+num2)
	case "-":
		fmt.Println("Результат: ", num1-num2)
	case "*":
		fmt.Println("Результат: ", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Результат: ", num1/num2)
		} else {
			fmt.Println("Ошибка: деление на ноль")
		}
	default:
		fmt.Println("Неверный оператор")
	}
}
