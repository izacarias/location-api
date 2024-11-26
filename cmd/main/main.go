package main

import "fmt"

// Defining a struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Creating an instance of the struct
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// Creating a pointer to the struct
	var pointer *Person
	pointer = &person

	// Accessing struct fields through the pointer
	fmt.Println("First Name:", pointer.FirstName)
	fmt.Println("Last Name:", pointer.LastName)
	fmt.Println("Age:", pointer.Age)
}
