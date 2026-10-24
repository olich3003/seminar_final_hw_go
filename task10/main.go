package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, i, number int
	var sum, max, min int
	var even, odd int
	var av float64
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
		fmt.Print("Введите целое число ", i, ": ")
		readCount, inputError = fmt.Scan(&inputText)
		number, inputError = strconv.Atoi(inputText)

		if readCount != 1 || inputError != nil {
			fmt.Println("Ошибка: нужно ввести целое число")
			return
		}

		if i == 1 {
			max = number
			min = number
		}

		sum = sum + number
		if number > max {
			max = number
		}

		if number < min {
			min = number
		}

		if number%2 == 0 {
			even = even + 1
		} else {
			odd = odd + 1
		}
	}

	av = float64(sum) / float64(n)
	fmt.Println("Сумма:", sum)
	fmt.Println("Среднее значение:", av)
	fmt.Println("Максимальное значение:", max)
	fmt.Println("Минимальное значение:", min)
	fmt.Println("Количество чётных чисел:", even)
	fmt.Println("Количество нечётных чисел:", odd)
}
