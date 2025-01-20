package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name     string
	Enroll   int
	Age      int
	IsPassed bool
}

func Json() {
	stud1 := Student{Name: "shivam", Enroll: 42, Age: 22, IsPassed: true}

	fmt.Println("Student1 details :", stud1)
	// Marshalling (Encoding Data into Json Format)
	jsonData, err := json.Marshal(stud1)

	if err != nil {
		fmt.Println("error on Marshaling Data:", err)
		return
	}
	fmt.Println("Json Data:-> \n", string(jsonData))

	// UnMarshalling (Decoding Data into Json Format)

	var StudData Student

	err = json.Unmarshal(jsonData, &StudData)

	if err != nil {
		fmt.Println("Error while UnMarshalling (Decoding) Data :", err)
		return
	}
	fmt.Println("UnMarshalled Data :", StudData)

}
