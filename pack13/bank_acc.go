package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (acc *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		acc.Balance += amount
		fmt.Printf("Пополнение: %.2f\n", amount)
	} else {
		fmt.Println("Сумма пополнения должна быть положительной")
	}
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Сумма снятия должна быть положительной")
	}
	if amount > acc.Balance {
		return errors.New("Недостаточно средств на счёте")
	}
	acc.Balance -= amount
	fmt.Printf("Снятие: %.2f\n", amount)
	return nil
}

func (acc BankAccount) CheckBalance() float64 {
	return acc.Balance
}

func main() {
	account := BankAccount{
		Owner:   "Иван Иванов",
		Balance: 1000,
	}

	fmt.Printf("Владелец: %s\n", account.Owner)
	fmt.Printf("Текущий баланс: %.2f\n", account.CheckBalance())

	account.Deposit(500)
	fmt.Printf("Баланс после пополнения: %.2f\n", account.CheckBalance())

	err := account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}

	err = account.Withdraw(300)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	}
	fmt.Printf("Баланс после снятия: %.2f\n", account.CheckBalance())
}
