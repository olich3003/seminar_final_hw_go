package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, result float64
	var operation string
	var inputTextA, inputTextB string
	var inputErrorA, inputErrorB error

	fmt.Print("Введите два числа через пробел: ")
	fmt.Scan(&inputTextA, &inputTextB)
	a, inputErrorA = strconv.ParseFloat(inputTextA, 64)
	b, inputErrorB = strconv.ParseFloat(inputTextB, 64)

	if inputErrorA != nil || inputErrorB != nil {
		fmt.Println("Ошибка: нужно ввести два числа")
		return
	}

	fmt.Print("Введите операцию +, -, *, /: ")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: делить на ноль нельзя")
			return
		}
		result = a / b
	default:
		fmt.Println("Ошибка: неизвестная операция")
		return
	}

	fmt.Println("Результат:", result)
}
