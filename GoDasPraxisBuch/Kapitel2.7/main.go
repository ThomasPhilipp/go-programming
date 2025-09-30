package main

import "fmt"

// Beispiel für einen eigenen Typ (Struktur) bzw. Typen

type user struct {
	firstName string
	lastName string
	address
}

type address struct {
	street string
	city string
}

func main() {
	// Declare struct as composite literals
	user1 := user {
		firstName: "Max",
		lastName: "Mustermann",
		address: address{
			street: "Musterstraße 1",
			city: "Musterstadt",
		},
	}

	fmt.Println("User1:", user1)
	fmt.Printf("User1: %+v\n", user1)

	// Declare anonymous struct
	user2 := struct {
		firstName string
		lastName string
		street string
		city string
	}{
		"Bob",
		"Andrew",
		"Backerystreet",
		"Lakeside",
	}
	fmt.Println("User2:", user2)
	fmt.Printf("User2: %+v\n", user2)
}