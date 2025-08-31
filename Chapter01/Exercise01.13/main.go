package main

import (
	"fmt"
	"time"
)

func main() {
	// Variant 1: declare a pointer using var --> initial value is nil
	var count1 *int

	// Variant 2: declare a point using new
	count2 := new(int)

	// Variant 3: declare a pointer from an existing variable
	countTemp := 5 // temporary variable
	count3 := &countTemp // & = address of

	// Variant 4: declare a pointer from a type (w/o temporary variable)
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time  : %#v\n", t)
}