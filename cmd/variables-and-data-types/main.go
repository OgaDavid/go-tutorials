package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Father *Person
}

func main() {
	// power := 9000
	// fmt.Printf("It's over %d\n", power)

	david := &Person{} // is a pointer with type *Person
	number := new(int)

	*number = 24

	fmt.Println(*number)

	john := Person{"John", 22, nil}

	david.Age = 30
	david.Name = "David"
	david.Father = &john
	fmt.Printf("%p\n", david)

	// structs()

	slices()

}
