package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, i int
	var number float64
	var pos, neg, zero int
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

	for i = 1; i <= n; i++ {
		fmt.Print("Введите число ", i, ": ")
		readCount, inputError = fmt.Scan(&inputText)
		number, inputError = strconv.ParseFloat(inputText, 64)
		if readCount != 1 || inputError != nil {
			fmt.Println("Ошибка: нужно ввести число")
			return
		}
		if number > 0 {
			pos = pos + 1
		} else if number < 0 {
			neg = neg + 1
		} else {
			zero = zero + 1
		}
	}

	fmt.Println("Положительных чисел:", pos)
	fmt.Println("Отрицательных чисел:", neg)
	fmt.Println("Нулей:", zero)
}
