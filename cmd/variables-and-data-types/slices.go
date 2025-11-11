package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func slices() {
	saiyans := []*Saiyan{
		{Name: "Goku", Power: 9000},
		{Name: "Vegeta", Power: 8000},
	}

	powers := extractPowers(saiyans)

	fmt.Println(powers)
}

func extractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))

	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}

	return powers
}
