package main

import "fmt"

// Beispiel für eigene Basis-Typen zurr Umwandlung von Meter in Zentimeter

type (
	meter int
	centimeter int
)

func main() {
	m := meter(5)
	fmt.Println("Input [m]:", m)
	fmt.Println("Output [cm]:", meterToCentimeter(m))
}

func meterToCentimeter(m meter) centimeter {
	return centimeter(m * 100);
}
