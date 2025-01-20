package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func FileRead() {
	// Open the file
	file, err := os.Open("new_file.txt")

	if err != nil {
		fmt.Println("error while opening file", err)
		return
	}
	// close resources
	defer file.Close()

	// create a buffer to read the file content
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break // END of File reached.
		}

		if err != nil {
			fmt.Println("Error while reading the file ", err)
			return
		}
		fmt.Println("value of n is :", n)
		// process the read content (in this case just print it)
		fmt.Println(string(buffer[:n]))

	}

	// Method-2 to Read file

	// Read the entire file in to the byte slice
	content, err := ioutil.ReadFile("new_file.txt")

	if err != nil {
		fmt.Println("error while reading the file data ", err)
		return
	}
	// Process the file content (in this example, just print it)

	fmt.Println(string(content))

}
