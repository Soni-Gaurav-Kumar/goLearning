package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

type Address struct {
	houseNo int
	street  string
	state   string
}

type Contact struct {
	email string
	phone string
}

type Employee struct {
	Person   // Embeded struct
	Address  // Embeded struct
	Contact  // Embeded struct
	Position string
}

func Structs() {
	// Method 1: Using the var keyword
	var person2 Person
	person2.firstName = "Alice"
	person2.lastName = "Smith"
	person2.age = 25
	fmt.Println("person2 Details: ", person2)

	// Method 2: Using a struct literal
	person3 := Person{
		firstName: "Bob",
		lastName:  "Johnson",
		age:       35,
	}
	fmt.Println("person3 Details: ", person3)

	// Method 3: Using the new keyword (returns a pointer to the struct)
	person4 := new(Person)
	person4.firstName = "Eve"
	person4.lastName = "Tayor"
	person4.age = 22
	fmt.Println("person4 Details: ", person4)

	// Creating an instance of the Employee struct (struct within struct) using Struct literals
	employee := Employee{
		Person: Person{
			firstName: "Gaurav",
			lastName:  "Soni",
			age:       22,
		},
		Address: Address{
			houseNo: 0,
			street:  "Jawahar Nagar",
			state:   "Uttar Pradesh",
		},
		Contact: Contact{
			email: "sonigaurav950@gmail.com",
			phone: "966XXXXXX",
		},
		Position: "SDE Intern",
	}

	fmt.Println("Employee Details: ", employee)

}
