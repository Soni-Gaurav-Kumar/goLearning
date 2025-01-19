package main

import "fmt"

func Ifelse() {
	// Example 1: if else-if else statement
	z := 7

	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but not 10")
	} else {
		fmt.Println("z is 5 or less")
	}

	// Example 4: Checking both conditions using &&
	x := 10
	y := 8

	if x > 5 && y < 10 {
		fmt.Println("Both conditions are true")
	} else {
		fmt.Println("At least one condition is false")
	}

}
