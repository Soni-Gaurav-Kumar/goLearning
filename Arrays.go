package main

import "fmt"

func Arrays() {
	// Method 1
	var names [5]string

	fmt.Println("names :", names)
	// names: [     ]                     output
	fmt.Printf("names : %q/n", names)
	// names: ["" "" "" "" ""]            output

	names[0] = "ankit"
	names[1] = "shivam"
	names[2] = "kishan"
	names[3] = "Anurag"

	fmt.Println("All Names in  names array :", names)
	fmt.Println("Name at index 1 :", names[1])

	//  Method 2 Initialize and assign values to an array in a single line
	numbers := [6]int{51, 23, 45, 67, 21}
	fmt.Println("numbers are: ", numbers)

	fmt.Println("Length of numbers array: ", len(numbers))
}

/*
NOTE: In Go when we declare an array or slice, the elements are initialized to their
 zero values. The zero values are default value assigned to variables of a certain type
 when no explicit value is provided

 For numeric types (int, float, etc.), the zero value is 0. For strings, it is an empty string
(""). For boolean types, it is false, and for pointers or complex types, it is nil.

*/
