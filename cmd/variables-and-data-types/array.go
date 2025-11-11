package main

import "fmt"

type Number int

func arrays() {
	numbers := []Number{1, 2, 3}

	scores := make([]int, 10, 20)

	fmt.Println(scores)

	fmt.Println(len(numbers))

	fmt.Println(numbers)

	numbers = append(numbers, 4)
	fmt.Println(len(numbers), cap(numbers))
}
