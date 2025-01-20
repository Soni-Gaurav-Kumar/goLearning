package main

import (
	"fmt"
	"io"
	"os"
)

func FileWrite() {
	file, err := os.Create("new_file.txt")

	if err != nil {
		fmt.Println("error while creating file", err)
		return
	}

	defer file.Close()

	initialContent := "Hello, I am learning File Handling in Go Language "

	// write the initial content to the file using io.WritrString

	bytes, errr := io.WriteString(file, initialContent+"\n")

	if errr != nil {
		fmt.Println("Issue while writing on file", errr)
		return
	}

	fmt.Printf("file created with initial content and have %d bytes\n", bytes)

}
