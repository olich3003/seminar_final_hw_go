package main

import (
	"fmt"
	"strconv"
)

func printTable(n float64) {
	var i int
	var res float64
	for i = 1; i <= 10; i++ {
		res = n * float64(i)
		fmt.Println(n, "*", i, "=", res)
	}
}

func main() {
	var n float64
	var inputText string
	var readCount int
	var inputError error

	fmt.Print("Введите число N: ")
	readCount, inputError = fmt.Scan(&inputText)
	n, inputError = strconv.ParseFloat(inputText, 64)
	if readCount != 1 || inputError != nil {
		fmt.Println("Ошибка: нужно ввести число")
		return
	}
	printTable(n)
}
