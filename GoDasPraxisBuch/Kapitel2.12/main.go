package main

import "fmt"

// Beispiele für Arrays und Slices

func main() {
	// Declare array (fix size, e.g. IP address of 4 bytes)
	fmt.Println("--- Array ---")
	ip := [4]byte{127,0,0,1}
	fmt.Println("Array:", ip)

	// Declare slice (flexible size)
	fmt.Println("--- slice ---")
	var s []int
	print(s)
	s = append(s, 1)
	print(s)
	s = append(s, 2)
	print(s)
	s = append(s, 3)
	print(s)
	s = append(s, 4)
	print(s)

	// Alternative approach to declare slice using make()
	fmt.Println("--- slice2 ---")
	n := []string{"Joe", "Bob", "Henry", "Ben", "Andrew"}
	for i /*index*/, v /*value*/ := range n {
		fmt.Printf("%02d: %s\n", i, v)
	}

	fmt.Println(n[0:3])
	fmt.Println(n[:3]) // same as before
	fmt.Println(n[3:5])
	fmt.Println(n[3:]) // same as before

	// Declare map
	fmt.Println("--- map ---")
	var m map[int]string = make(map[int]string)
	m[1] = "Joe"
	m[2] = "Bob"
	m[3] = "Henry"
	m[4] = "Ben"
	m[5] = "Andrew"
	fmt.Println(m)
}

func print(s []int) {
	fmt.Printf("%p - len: %d cap: %d %v\n", s, len(s), cap(s), s)
}