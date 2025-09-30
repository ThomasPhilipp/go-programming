package main

import "fmt"

// Beispiel für einen Pointer

func main() {
	fmt.Println("--- simpleUsage ---")
	simpleUsage()

	fmt.Println("--- functionUsage ---")
	functionUsage()
}

func simpleUsage() {
	var num int
	var numPointer *int // Type: Pointer to int

	num = 7
	numPointer = &num // Get memory address "&"

	fmt.Println("Initial value: ", num, numPointer)

	*numPointer = 14 // Read or change value of a variable -> "*" aka dereference
	fmt.Println("Changed value: ", num /*value*/, numPointer /*pointer address*/)
}

func functionUsage() {
	foo := foo()
	fmt.Println("Initial value: ", foo /*pointer address*/, *foo /*value*/)
}

// Returns a pointer
func foo() *int {
	a := 111
	return &a
}