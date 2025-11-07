package main

import "fmt"

// PET STORE SYSTEM
// What can ALL pets do, regardless of what type they are?

// Interface = "Any pet must be able to do these things"
type Pet interface {
	Speak() string
	Eat(food string) string
	Play() string
}

// CREATE STRUCTS (actual data about each pet)

// Dog has a Name and Breed
type Dog struct {
	Name  string
	Breed string
}

// Cat has a Name and Color
type Cat struct {
	Name  string
	Color string
}

// ATTACH METHODS (how each pet does it differently)

// How does a Dog Speak?
func (d Dog) Speak() string { // we are also attaching the method to the Dog struct
	return d.Name + " says Woof!"
}

// How does a Dog Eat?
func (d Dog) Eat(food string) string {
	return d.Name + " is chomping on " + food
}

// How does a Dog Play?
func (d Dog) Play() string {
	return d.Name + " is fetching the ball!"
}

// How does a Cat Speak?
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// How does a Cat Eat?
func (c Cat) Eat(food string) string {
	return c.Name + " is nibbling on " + food
}

// How does a Cat Play?
func (c Cat) Play() string {
	return c.Name + " is chasing the laser pointer!"
}

func interfaces() {
	// CREATE INSTANCES OF PETS
	dog := Dog{Name: "Buddy", Breed: "Golden Retriever"}
	cat := Cat{Name: "Whiskers", Color: "Tabby"}

	// CREATE A SLICE OF PETS (I can store different types that implement the Pet interface)
	pets := []Pet{dog, cat}

	// INTERACT WITH EACH PET
	for _, pet := range pets {
		fmt.Println(pet.Speak())
		fmt.Println(pet.Eat("treats"))
		fmt.Println(pet.Play())
		fmt.Println()
	}
}
