package main

import "fmt"

// NOTE : Go doesn't Support Nested Functions

func simpleFunction() {
	fmt.Println("Below are the examples of functions (here sum, product are calculate using functions)")
	fmt.Println("This is a simple function.")
}
func add(a, b int) int {
	return a + b
}

func multiply(x, y int) (result int) {
	result = x * y
	return
}

func Functions() {
	simpleFunction()
	sum := add(10, 5)
	fmt.Println("sum of two Numbers: ", sum)
	product := multiply(10, 5)
	fmt.Println("product of two Numbers: ", product)

}

// encorrect ways of define function
// func abc()
// {

// }

// func pqr(a int, b){

// }
