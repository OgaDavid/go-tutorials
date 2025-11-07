package main

import "fmt"

type Human struct {
	Name    string
	Age     int
	Country string
}

func (p *Human) Greet() {
	p.Name = "Samuel"
	// fmt.Println("Hello, my name is", p.Name)
}

func structs() {

	student := Human{
		Name:    "David",
		Age:     25,
		Country: "USA",
	}

	student.Greet()
	fmt.Println("After Greet method, Name is:", student.Name)

	// structs + slices
	people := []Human{
		{Name: "Alice", Age: 30, Country: "Canada"},
		{Name: "Bob", Age: 22, Country: "UK"},
	}

	fmt.Printf("Student: %+v\n", student)
	fmt.Printf("People: %+v\n", people)

}
