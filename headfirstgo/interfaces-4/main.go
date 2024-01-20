package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

func (d Dog) Scoobie() string {
	return "Rut Row"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

// When you have a value of a concrete type assigned to a variable with a interface type,
// a type asseriton lets you get the concrete type back.
func main() {
	animals := []Animal{Dog{}, Cat{}}

	for _, animal := range animals {
		// Type assertion to access the underlying type
		// Since we are back to the concrete type we have access to all the Methods too
		if dog, ok := animal.(Dog); ok {
			fmt.Println("It's a dog:", dog.Speak())
			fmt.Println("Scoobie:", dog.Scoobie())
		} else if cat, ok := animal.(Cat); ok {
			fmt.Println("It's a cat:", cat.Speak())
		}

		// Comma-Ok Idiom:
		// if variable, ok := someExpression.(SomeType); ok {...}
	}

	for index, pet := range animals {
		fmt.Printf("Pet Num# %d is a %s\n", index, pet.Speak())
	}
}