package main

import "fmt"

func Defers() {
	fmt.Println("Start of the Program")
	// The function inside defer will be executed when the surrounding function (main in this case) exists

	defer fmt.Println("This will be executed at the end")
	defer fmt.Println("This will be executed second (after middle)")
	defer fmt.Println("This will be executed first (after middle)")

	fmt.Println("Middle of the Program")

	/* Note: When you have multiple defer statements in a function, they are executed in a last-in,
	first-out (LIFO) order. The last defer statement you encounter in the code will be the
	first one to be executed when the function exits.
	*/
}
