package main

import (
	"fmt"
	"io"
	"net/http"
)

func WebRequests() {
	// Make a Get Requests

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/49")

	// check data type of response we get
	fmt.Printf("type of res : %T\n", res)

	if err != nil {
		fmt.Println("error on making GET request :", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("error on reading data from req Body :", err)
		return
	}

	fmt.Println("data :", string(data))
}
