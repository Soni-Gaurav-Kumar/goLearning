package main

import (
	"fmt"
	"strconv"
)

func DataTypeConv() {
	// int to float conversion
	var integerNum int = 42
	var floatNum float64 = float64(integerNum)
	fmt.Println("floatNum :", floatNum)

	// String conversion
	var number int = 142
	str := strconv.Itoa(number) // Integer to string

	fmt.Println("str :", str)

	strNum := "123"
	num, err := strconv.Atoi(strNum)

	fmt.Println("num :", num)
	if err != nil {
		fmt.Println("Error: not a valid integer while converting from string to int")
	}

	// Convert String to Float
	strFloat := "3.14"
	numFloat, err := strconv.ParseFloat(strFloat, 64)

	if err == nil {
		fmt.Println("numFloat :", numFloat)
	}
}
