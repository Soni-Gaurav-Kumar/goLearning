package main

import "fmt"

func ForLoop() {
	// Example 1: Basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Example 2: Infinite loop with break statement (working as while loop in other languages)
	counter := 0

	for {
		fmt.Printf("Loop count  - %d\n", counter)
		counter++
		if counter == 5 {
			break
		}
	}

	numbers := []int{67, 45, 32, 87, 98, 54, 49}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	message := "Hello, Golang!"
	// Looping over characters of a string using range
	for index, char := range message {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

}
