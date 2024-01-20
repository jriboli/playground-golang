package main

import "fmt"

// Defining Methods for a Struct
type MyType string 

// Go uses receiver parameters instead of "self" or "this" seen in other languages
// Syntacs of a method:
// func ({Reciever Parameter Name} {Receiver Parameter Type}) methodName(paramName paramType) returnType {..}
func (m MyType) sayHi() {
	fmt.Println("Hi")
}

//Pointer receiver parameters
type Number int 

// We have to pass the receiver in as a point if want it to update the value outside of the func
func (n *Number) double() {
	// What is the value of n at this point - N is a pointer for the address of Number
	fmt.Println(n)
	// What is the value of *n at this point - N we get the value at the pointer
	fmt.Println(*n)
	// What is the value of &n at this point - This is a pointer to the pointer of the address of Number
	fmt.Println(&n)
	*n *= 2
	fmt.Println(n)
}

type Person struct {
	Name string
	Age int
}