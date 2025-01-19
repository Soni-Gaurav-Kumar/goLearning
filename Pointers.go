package main

import "fmt"

func ModifyValueByReference(num *int) {
	*num = *num + 10
}

func Pointers() {
	// Declare a variable and a pointer
	var num int = 2
	var ptr *int

	// Initialize the pointer with the address of the variable
	ptr = &num

	// Access the value through the pointer
	fmt.Println("pointer ptr contain :", ptr)
	fmt.Println("Value of num :", num)
	fmt.Println("Value of num through ptr :", *ptr)

	// variable declaration and assignment on the same line.

	data := "soni"
	ptr1 := &data
	fmt.Println("Value through pointer: ", *ptr1)

	/* NOTE :In Go, pointers are initialized with nil by default if not explicitly set to point to a
	valid memory address. A nil pointer doesn't point to any valid memory location.
	*/

	var ptr2 *int
	// Check if the pointer ptr2 is nil
	if ptr2 == nil {
		fmt.Println("Pointer ptr2 is nil")
	}

	val := 43
	fmt.Println("Before modify by ref :", val)

	ModifyValueByReference(&val)
	fmt.Println("After modify by refe :", val)

}
