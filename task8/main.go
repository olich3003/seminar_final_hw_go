package main

import (
	"fmt"
	"strconv"
)

func isPrime(n int) bool {
	var div int

	if n < 2 {
		return false
	}

	for div = 2; div <= n/div; div++ {
		if n%div == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	var pr bool
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

	pr = isPrime(n)
	if pr {
		fmt.Println("Число простое")
	} else {
		fmt.Println("Число не является простым")
	}
}
