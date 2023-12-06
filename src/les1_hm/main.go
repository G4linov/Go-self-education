package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Математические операции
	var a = 1
	var b = 3
	var c = b - a
	fmt.Println(c)
	c = b * a
	fmt.Println(c)
	a++
	b--
	c = a % b
	fmt.Println(a, b, c)

	// Преобразование типов

	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

	fmt.Println(s, i)

	// Домашка пункт 4

	var code = "104"
	var number = 35

	codeint, _ := strconv.Atoi(code)
	numberstr := strconv.Itoa(number)

	fmt.Println(codeint, numberstr)

}
