package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, i int
	var number, maximum, minimum float64
	var inputText string
	var readCount int
	var inputError error

	fmt.Print("Введите количество чисел N: ")
	readCount, inputError = fmt.Scan(&inputText)
	n, inputError = strconv.Atoi(inputText)

	if readCount != 1 || inputError != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}

	if n <= 0 {
		fmt.Println("Ошибка: количество чисел должно быть больше нуля")
		return
	}

	for i = 1; i <= n; i++ {
		fmt.Print("Введите число ", i, ": ")
		readCount, inputError = fmt.Scan(&inputText)
		number, inputError = strconv.ParseFloat(inputText, 64)
		if readCount != 1 || inputError != nil {
			fmt.Println("Ошибка: нужно ввести число")
			return
		}
		if i == 1 {
			maximum = number
			minimum = number
		}
		if number > maximum {
			maximum = number
		}
		if number < minimum {
			minimum = number
		}
	}

	fmt.Println("Максимальное значение:", maximum)
	fmt.Println("Минимальное значение:", minimum)
}
