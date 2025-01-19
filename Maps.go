package main

import "fmt"

func Maps() {
	// Method 1- Creating a Map
	studentGrades := make(map[string]int)

	// Adding Key Value pairs
	studentGrades["Alice"] = 76
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 88

	// Accessing values
	fmt.Println("Alice's Grade:", studentGrades["Alice"])

	// Modifying values
	studentGrades["Bob"] = 88

	// Deleting a key-value pair
	delete(studentGrades, "Charlie")

	// Checking if a Key exists
	grade, exists := studentGrades["David"]
	fmt.Println("David Grade Exists :", exists)
	fmt.Println("David Grade :", grade)

	// Iterating Over Map
	fmt.Println("\nStudent Grades: ")
	for name, grade := range studentGrades {
		fmt.Printf("%s : %d\n", name, grade)
	}

	// Method 2:- Creating a map using a literal (make like dict in other language)

	fruitPrices := map[string]int{
		"Apple":  120,
		"Banana": 80,
		"Orange": 60,
		"Mango":  180,
	}

	fmt.Println("\nfruits : ", fruitPrices)

}
