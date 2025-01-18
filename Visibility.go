package main

import "fmt"

func Visibility() {
	// private member initialised with first letter 'small' (unexported) can be visible within same package
	var privateVariable int = 10

	// public members/function will be initiallised with first letter 'Capital' (exported) and can be visible from other package
	var PublicVariable int = 49

	fmt.Println("privateVariable :", privateVariable)
	fmt.Println("PublicVariable :", PublicVariable)

	// same will be applied for functions also

}
