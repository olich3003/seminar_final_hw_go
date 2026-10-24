package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, i, sum int
	var average float64
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

	if n <= 0 {
		fmt.Println("Ошибка: N должно быть больше нуля")
		return
	}

	for i = 1; i <= n; i++ {
		sum = sum + i
	}

	average = float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее значение:", average)
}
