package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

func PerformGetRequest(myUrl string) {

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetched Data :", response.Status)
		return
	}

	var todo Todo

	err = json.NewDecoder(bytes.NewReader(body)).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding Response in first :", err)
		return
	}
	fmt.Println("Todo ID :", todo.ID)
	fmt.Println("Todo Title :", todo.Title)
	fmt.Println("Todo UserID :", todo.UserID)
	fmt.Println("Todo Completed :", todo.Completed)

	// another way to decode data

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var todo1 Todo
	err = json.Unmarshal(body, &todo1)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

}

func performPostRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/"
	todo := Todo{
		UserID:    43,
		ID:        2,
		Title:     "BigBasket",
		Completed: true,
	}

	// Convert todo sturct into Json (marshalling)

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	// Convert the JSON byte slice to a string
	jsonStr := string(jsonData)

	// Create an io.Reader from the string
	jsonReader := strings.NewReader(jsonStr)

	// Send the POST request
	resp, err := http.Post(myUrl, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status code :", resp.Status)

}

func performUpdateRequest(myUrl string) {
	todo := Todo{
		UserID:    43,
		ID:        2,
		Title:     "Update Todo",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a PUT request
	request, err := http.NewRequest(http.MethodPut, myUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	request.Header.Set("content-type", "application/json")

	// Send the Request
	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error in sending PUT request :", err)
		return
	}

	defer response.Body.Close()
	fmt.Println("Response Status:", response.Status)

}

func performDeleteRequest(myUrl string) {
	// Create a Delete Request

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)

	if err != nil {
		fmt.Println("Error in creating Delete request:", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error in sending Delete request :", err)
		return
	}

	defer res.Body.Close()
	fmt.Println("Response Status:", res.Status)

}

func CRUD() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	PerformGetRequest(myUrl)
	performPostRequest()
	performUpdateRequest(myUrl)
	performDeleteRequest(myUrl)

}
