package main

import "fmt"

func Variables() {
	var age int = 22
	var height float64 = 5.3
	var name string = "Gaurav"
	var isAdult bool = true
	const pi float64 = 3.14159

	// short Hand for declaring variables

	team := "uplifters"
	college := "NIT Srinagar"
	members := 18

	fmt.Println("age:", age)
	fmt.Println("height:", height)
	fmt.Println("name:", name)
	fmt.Println("Is Adult:", isAdult)
	fmt.Println("value of Pi:", pi)

	fmt.Println("Team Name: ", team)
	fmt.Println("college Name: ", college)
	fmt.Println("No. of Members: ", members)

}
