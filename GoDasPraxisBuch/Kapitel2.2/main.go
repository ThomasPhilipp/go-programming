package main

import (
	"fmt"
	"os"
)

// Beispiele für die Gruppierung von Variablen

var (
	home = os.Getenv("HOME")
	user = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

// Beispiel für eine Konstante mit einem automatischen Zähler
const (
	_ // no assignment
	MON = iota // automatic counter
	TUE
	WED
	THI
	FRI
	SAT
	SUN
)

func main() {
	fmt.Println("HOME variable: ", home)
	fmt.Println("USER variable: ", user)
	fmt.Println("GOPATH variable: ", gopath)

	// Kurzdeklaration (Definition des Typs erfolg über den Wert der Variablen)
	a, b, c := 0, 0, 7
	fmt.Println("Erkennung: ", a, b, c)
}