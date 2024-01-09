package main

import (
	"fmt"
	"strings"
)

func main() {
	// [CONSTANT TYPE]
	// GoLang will not allow you to change
	const conferenceTickets = 50

	// [INFERRED TYPE]
	// GoLang will infer the data type by the assistment
	var conferenceName = "Go Conference"
	var remainingTickets int = 50

	// [DECLARED TYPE]
	// This can be overwritten or more so declared
	// example:
	var someNumber int = 0
	fmt.Println("Number is", someNumber)
	// Like setting type of uint - can not be negative
	var somePositiveNumber uint = 0
	// Wont let us set it to a negative
	// somePositiveNumber = -1 -- Will throw an error
	fmt.Println("Another number", somePositiveNumber)

	// [SYNTACTITCAL SUGAR]
	// dont add 'var
	// dont add datatype
	// Can not be used with constants
	custName := "Rocket Raccoon"
	fmt.Println("Name is", custName)

	// [ARRAY]
	// Ways of declaring/initializing the array
	var names [50]string
	names[0] = "Gamora"

	// Array will start values
	var bookings = [50]string{"Rocket", "Drax"}
	fmt.Println(bookings)

	// SLICE - non fixed array
	var bookingSlice []string
	// or var bvookingSlice = []string{}
	bookingSlice = append(bookingSlice, "Rocket Raccoon")

	greetUsers(conferenceTickets, remainingTickets, conferenceName)

	// [LOOP]
	for {
		fmt.Printf("We have a total of %v and there are still %v left\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")

		// Declaritive Type
		var firstName string
		var lastName string
		var userTickets int

		// ENTER NAME
		for {
			fmt.Println("Enter your first name:")
			fmt.Scan(&firstName)
			fmt.Println("Enter your last name:")
			fmt.Scan(&lastName)

			if len(firstName) < 2 || len(lastName) < 2 {
				fmt.Println("Invalid name, please use atleast 2 characters")
			} else {
				break
			}
		}

		for {
			fmt.Println("How many tickets do you need:")
			fmt.Scan(&userTickets)
			if userTickets < 0 {
				fmt.Println("Invalid number, must be positive")
			} else if remainingTickets-userTickets < 0 {
				fmt.Println("You are trying to purchase more than we have, only have", remainingTickets, "left")
			} else {
				break
			}
		}

		remainingTickets -= userTickets
		fmt.Printf("Thank you %v %v for booking %v tickets.\n",
			firstName, lastName, userTickets)

		bookingSlice = append(bookingSlice, firstName+" "+lastName)
		// DEBUGGING VALUES
		// fmt.Printf("The whole array: %v\n", bookingSlice)
		// fmt.Printf("The first value: %v\n", bookingSlice[0])
		// fmt.Printf("Array type: %T\n", bookingSlice)
		// fmt.Printf("Array length: %v\n", len(bookingSlice))

		printFirstNames(bookingSlice)

		fmt.Printf("These are all our bookings: %v\n", bookingSlice)

		noTicketsRemaining := remainingTickets <= 0
		// [IF/ELSE BLOCKS]
		// Conditional Statement
		if noTicketsRemaining {
			// End the program
			fmt.Println("Our conference is booked. Come back next year.")
			break
		}
	}

	// DO WHILE LOOP
	// Similar to IF/ELSE statement
	var loopCondition bool = true
	var loopCounter = 0
	for loopCondition {
		if loopCounter > 10 {
			loopCondition = false
		}
		fmt.Println("WHILE LOOP, count - ", loopCounter)
		loopCounter++
	}
}

func greetUsers(confTickets int, remainTickets int, confName string) {
	// Print the type of a variable
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", confTickets,
		remainTickets, confName)

	// String concatenate
	fmt.Println("---------------------------------------")
	fmt.Println("----", confName, "Application----")
	fmt.Println("---------------------------------------")
	// String format
	fmt.Printf("Welcome to the %v booking application\n", confName)

}

func printFirstNames(bookingSlice []string) {

	firstNames := []string{}
	for _, booking := range bookingSlice {
		// _ is a blank identifier
		// used to hold a variable defined here that we are not using
		var names = strings.Fields(booking)
		//var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("Only the first names: %v\n", firstNames)
}

func getUserName() {

}
