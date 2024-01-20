package main

import "fmt"

type Printer interface {
	Print()
}

type HewletPacker struct {

}

// HewletPacker satifies the Printer interface with
func (h HewletPacker) Print() {
	fmt.Println("This is the HewletPacker")
}

func main() {
	// An interface variable - variable that is of type Printer Interface
	// Can hold any value that implements Printer
	var printer Printer
	printer = HewletPacker{}

	printer.Print()
}