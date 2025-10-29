package main

import "fmt"

// array and slices
func printArray() {

numbers := []int{1, 2, 3}
fmt.Println(len(numbers), cap(numbers)) // 3, 3

numbers = append(numbers, 4)
fmt.Println(len(numbers), cap(numbers)) // 4, 6

	
}