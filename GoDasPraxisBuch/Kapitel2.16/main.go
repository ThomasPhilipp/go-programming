package main

import "fmt"

func main() {
	fmt.Println("--- switch1 ---")
	switch1(7)

	fmt.Println("--- switch2 ---")
	switch2(3)

	fmt.Println("--- switch3 ---")
	switch3(-1)

	fmt.Println("--- switch4 ---")
	switch4(25)

	fmt.Println("--- switch5 ---")
	switch5()
}

// Simple usage
func switch1(value int) {
	switch value {
	case 0:
		fmt.Println("Value equals 0")
	case 1-9: // or 1, 2, 3, 4, 5, 6, 7, 8, 9
		fmt.Println("Value below 10")
	default:
		fmt.Println("Value greater 10")
	}
}

// Variable declaration
func switch2(value int) {
	switch i:=value-1; i {
	case 0: fmt.Println("Value equals 0")
	default: fmt.Println("Value positive")
	}
}

// Switch as if statement
func switch3(value int) {
	switch {
	case value < 0: fmt.Println("Value negative")
	case value == 0: fmt.Println("Value equals 0")
	case value > 0: fmt.Println("Value positive")
	}
}

// fallthrough
func switch4(value int) {
	switch {
	case value > 0:
		fmt.Println("Value positive")
		fallthrough
	case value > 10:
		fmt.Println("Value greater 10")
		fallthrough
	case value > 20:
		fmt.Println("Value greater 20")
		fallthrough
	case value > 30:
		fmt.Println("Value greater 30")
	}
}

// type switch
func switch5() {
	var x interface{}

	switch x.(type) {
	case int:
		fmt.Println("Given value is of type int")
	case float64:
		fmt.Println("Given value is of type float64")
	default:
		fmt.Println("Given value is unknown:", x)
	}
}