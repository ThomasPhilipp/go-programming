package main

import "fmt"

// Beispiel für einen Pointer

func main() {
	var number int
	var pointer *int // Type: Pointer to int

	number = 7
	pointer = &number // Get memory address "&"

	fmt.Println("Initial: ", number, pointer)

	*pointer = 14 // Read or change value of a variable -> "*" aka dereference
	fmt.Println("Changed: ", number, pointer)
}