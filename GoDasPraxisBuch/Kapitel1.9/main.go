package main

import (
	"fmt"
	"os"
)

// Beispiele des "fmt"-Paket, das importiert wird. Dadurch können wir alle
// Funktionen, Typen, Variablen und Konstanten des Pakets ansprechen.

func main() {
	fmt.Print("Hallo ", "Print()\n")

	fmt.Println("Hallo ", "Println()")

	var s = "Printf()"
	fmt.Printf("Hallo, %s\n\n", s)

	// Erweiterte formartierte Ausgabe

	// %v: Ausgabe eines Werts einer beliebigen Variable
	// $+v: Ausgabe der Struktur und die Namen der einzelnen Felder
	// %#v: Ausgabe wie im Go-Code definiert
	// %T: Ausgabe des Types einer Variablen
	// %b: Ausgabe zur Basis 2 (binäre Zahl)
	// %d: Ausgabe zur Basis 10
	// %03d: Ausgabe wird mit 0 aufgefüllt
	// %p: Ausgabe der Speicheradresse, eines Pointers
	// %s: Ausgabe eines Strings
	// %x: Ausgabe eines Strings als hex-Wert
	var nr = 7
	var name = "James Bond"
	fmt.Printf("%03d: Hallo, %s\n", nr, name)

	// Erzeugen von Strings
	// fmt.Sprint(), fmt.Sprintln() und fmt.Sprintf()

	// Schreiben in ein Datenziel (io.Writer), zB: Datei, Standard Output, Webserver, etc.
	// fmt.Fprint(), fmt.Fprintln() und fmt.Fprintf()

	file, _ := os.Create("sample-file.txt")
	fmt.Fprintf(file, "%03d: Hallo, %s\n", nr, name)
	file.Close()
}