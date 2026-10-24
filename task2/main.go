package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
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

	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю")
	}

	if n%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечётное")
	}

	if n%3 == 0 {
		fmt.Println("Число делится на 3")
	} else {
		fmt.Println("Число не делится на 3")
	}

	if n%5 == 0 {
		fmt.Println("Число делится на 5")
	} else {
		fmt.Println("Число не делится на 5")
	}
}
