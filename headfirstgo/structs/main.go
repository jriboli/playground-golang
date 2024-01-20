package main

import (
	"fmt"
	"log"

	"headfirstgo/structs/calendar"
)

func main() {
	event := calendar.Event{}
	event.SetYear(2023)
	event.SetTitle("Someones Birthday")

	fmt.Println(event)

	date := calendar.Date{} // Create a Date
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal("Error occurred: ", err)
		// If value is invalid, log the error and exit
	}

	err = date.SetMonth(5) 
	if err != nil {
		log.Fatal("Error occurred: ", err)
	}

	err = date.SetDay(11) 
	if err != nil {
		log.Fatal("Error occurred: ", err)
	}

	fmt.Println(date)
	// This will no longer work as Date is in a different package
	//fmt.Println(date.year)
	fmt.Println(date.Year())

	value := MyType("a MyType value")
	value.sayHi()

	number := Number(2)
	number.double()
	fmt.Println("Number after doubling: ", number)

	// Some pointer knowledge 
	//Declare a variable
	var num int = 42

	// Declare a pointer and assign the address of the variable
	var ptr *int = &num

	//Access the value through the pointer
	fmt.Println("Value:", *ptr) // Output: 42

	// Pass the pointer to the function to update the value
	updateValue(ptr)
	fmt.Println("Value Updated:", *ptr) // Output: 84

	// Alternative way to call it
	updateValue(&num)
	fmt.Println("Value Again:", *ptr)

	// Using the 'new' fuction
	p := new(int) // p is a pointer to a zeroed int
	fmt.Println("Using New:", *p)

	person := Person{Name: "Rocket", Age: 32}
	fmt.Println("Current Age:", person.Age)

	// Add a year
	updatePerson(&person)
	fmt.Println("Add a year:", person.Age)
}

func updateValue(ptr *int) {
	*ptr *= 2
}

func updatePerson(p *Person) {
	p.Age++
}