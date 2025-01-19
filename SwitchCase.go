package main

import "fmt"

func SwitchCase() {
	// Example 1: Basic switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	// Example 2: Switch statement with multiple values
	month := "January"

	switch month {
	case "Janiary", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Summer")
	case "July", "August", "September":
		fmt.Println("Rainy")
	default:
		fmt.Println("Other Season")

	}

	// Example 3: Switch with expression
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
