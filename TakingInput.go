package main

import (
	"bufio"
	"fmt"
	"os"
)

func TakingInput() {
	var fname string
	fmt.Println("Enter Your First name (without Space): ")
	fmt.Scan(&fname)

	fmt.Printf("Hello, %s!\n", fname)

	var fullName string
	fmt.Println("Enter your Full Name (with space): ")
	reader := bufio.NewReader(os.Stdin)
	fullName, _ = reader.ReadString('\n')

	fmt.Printf("ohh, your full name is %s\n", fullName)

}
