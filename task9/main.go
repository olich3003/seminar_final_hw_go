package main

import (
	"fmt"
	"strconv"
)

func main() {
	var balance, sum float64
	var action, inputText string
	var readCount int
	var inputError error

	fmt.Print("Введите баланс: ")

	readCount, inputError = fmt.Scan(&inputText)
	balance, inputError = strconv.ParseFloat(inputText, 64)

	if readCount != 1 || inputError != nil {
		fmt.Println("Ошибка: баланс должен быть числом")
		return
	}

	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")
	fmt.Print("Введите номер действия: ")

	fmt.Scan(&action)
	if action == "2" || action == "3" {
		fmt.Print("Введите сумму: ")
		readCount, inputError = fmt.Scan(&inputText)
		sum, inputError = strconv.ParseFloat(inputText, 64)

		if readCount != 1 || inputError != nil {
			fmt.Println("Ошибка: сумма должна быть числом")
			return
		}
	}

	switch action {
	case "1":
		fmt.Printf("Баланс: %.2f\n", balance)

	case "2":
		balance = balance + sum
		fmt.Printf("Новый баланс: %.2f\n", balance)

	case "3":
		if sum > balance {
			fmt.Println("Ошибка: недостаточно средств на балансе")
			return
		}
		balance = balance - sum
		fmt.Printf("Новый баланс: %.2f\n", balance)

	default:
		fmt.Println("Ошибка: такого действия нет")
	}
}
