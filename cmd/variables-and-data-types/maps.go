package main

import "fmt"

func maps() {

	initializedMap := make(map[int]string)
	initializedMap[0] = "David"

	myMap := map[string]string{
		"name":     "Alice",
		"city":     "Wonderland",
		"country":  "England",
		"language": "English",
		"age":      "50",
	}

	myMap["profession"] = "Adventurer"
	myMap["age"] = "40"
	delete(myMap, "city")

	fmt.Println(myMap)
	fmt.Println("Map length:", len(myMap))
	fmt.Println("Initialized map:", initializedMap[0])
}
