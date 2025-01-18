package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("warning! denominator must not be zero")
	}

	return a / b, nil
}

func ErrorHandling() {
	// division, _ := divide(4, 0)   for egnoring variables we can use _ know as black identifier
	division, errMsg := divide(4, 0)
	fmt.Println("Division of two numbers: ", division)
	if errMsg != nil {
		fmt.Println("Error is: ", errMsg)
	}
}
