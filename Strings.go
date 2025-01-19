package main

import (
	"fmt"
	"strings"
)

func Strings() {
	// Splitting string based on some delimeter
	str := "apple,orange,banana"
	parts := strings.Split(str, ",")
	fmt.Println(parts) // Output: [apple orange banana]

	// Counting Occurences
	sentence := "one two three four two two five"
	count := strings.Count(sentence, "two")
	fmt.Println("Count:", count)

	//Trimming Whitespace:
	sentence2 := "      Hello, Go!                          "
	trimmed := strings.TrimSpace(sentence2)
	fmt.Println("Trimmed:", trimmed)

	// Concatenating Strings:

	str1 := "Hello World"
	str2 := "with Go!"

	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)
}
