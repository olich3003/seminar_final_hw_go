package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, i, sum int
	var inputText string
	var readCount int
	var inputError error

	fmt.Print("Введите целое число N: ")

	readCount, inputError = fmt.Scan(&inputText)
	n, inputError = strconv.Atoi(inputText)

	if readCount != 1 || inputError != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}

	for i = 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum = sum + i
		}
	}

	fmt.Println("Сумма подходящих чисел:", sum)
}
