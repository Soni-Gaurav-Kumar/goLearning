package main

import "fmt"

func Slices() {
	// Method 1:   Declare and initialize a slice
	numbers := []int{1, 2, 3, 4, 5}

	// Access and print slice elements and its length, capacity

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// Method 2:  Creating a slice with make: length = 3, capacity = 5
	nums := make([]int, 3, 5)
	fmt.Println("Slice:", nums)
	fmt.Println("Length:", len(nums), "Capacity:", cap(nums))

	// Method 3: // Creating an empty slice of strings with make
	var days = make([]string, 0)
	fmt.Println(days)

	// NOTE:  You can’t initialize the slice like this :-->>    var numbers = []string
}
