package main

import "fmt"

// Beispiele für erweiterte Funktionsaufrufe

// First-Class Citizens: ein Funktion = ein Typ -> Zuerweisung einer Variablen
type filterFunc func(string)bool

func main() {
	fmt.Println("--- swapNumbers ---")
	a := 3
	b := 7
	fmt.Println("Input:", a, b)
	b2, a2 := swapNumbers(a, b)
	fmt.Println("Output:", b2, a2)

	fmt.Println("--- sum ---")
	fmt.Println("Input:", 3, 7, 9)
	fmt.Println("Output:", sum(3, 7, 9))

	fmt.Println("--- filter ---")
	values := []string{"a", "aa", "aaa", "aaaa", "aaaaaa"}
	filterSample := func(v string) bool {
		if len(v) < 3 {
			return false
		}
		return true
	}
	fmt.Println("Filtered output:", filter(values, filterSample))

	fmt.Println("--- filterByLen ---")
	fmt.Println("Filtered output by len=1:", filter(values, filterByLen(1)))
	fmt.Println("Filtered output by len=2:", filter(values, filterByLen(2)))
	fmt.Println("Filtered output by len=3:", filter(values, filterByLen(3)))
	fmt.Println("Filtered output by len=4:", filter(values, filterByLen(4)))
	fmt.Println("Filtered output by len=5:", filter(values, filterByLen(5)))

	fmt.Println("--- defers1 ---")
	defers1()

	fmt.Println("--- defers2 ---")
	defers2()
}

// Mehrere Input and Return Werte
func swapNumbers(a, b int) (int, int) {
	return b, a
}

// Varadische Funktion
func sum(n...int /*slice = []int*/) int {
	sum := 0
	for _, v := range n {
		sum = sum + v
	}
	return sum
}

// First-Class Citizen
func filter(values []string, filter filterFunc) []string {
	var out []string
	for _, v := range values {
		if (filter(v)) {
			out = append(out, v)
		}
	}
	return out
}

// Closures
func filterByLen(length int) filterFunc {
	// Anonymous function with matching signature
	return func(value string) bool {
		if (len(value) < length) {
			return false
		}
		return true
	}
}

func defers1() {
	defer fmt.Println("1")
	defer fmt.Println("2") // wird vor dem Anderen defer ausgeführt (!)
	fmt.Println("3")
}

func defers2() {
	// Anonymous function
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
	}()
	fmt.Println("3")
}