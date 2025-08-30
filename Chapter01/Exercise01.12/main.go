package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	fmt.Printf("Count   (#v) : %#v \n", count)
	fmt.Printf("Count   (+v) : %+v \n", count)
	fmt.Printf("Count   (T)  : %T \n", count)
	fmt.Printf("Count   (d)  : %d \n", count)
	var discount float64
	fmt.Printf("Discount (#v): %#v \n", discount)
	fmt.Printf("Discount (+v): %+v \n", discount)
	fmt.Printf("Discount (T) : %T \n", discount)
	fmt.Printf("Discount (d) : %d \n", discount)
	var debug bool
	fmt.Printf("Debug  (#v)  : %#v \n", debug)
	fmt.Printf("Debug  (+v)  : %+v \n", debug)
	fmt.Printf("Debug  (T)   : %T \n", debug)
	var message string
	fmt.Printf("Message (#v) : %#v \n", message)
	fmt.Printf("Message (+v) : %+v \n", message)
	fmt.Printf("Message (T)  : %T \n", message)
	fmt.Printf("Message (s)  : %s \n", message)
	var emails []string
	fmt.Printf("Emails (#v)  : %#v \n", emails)
	fmt.Printf("Emails (+v)  : %+v \n", emails)
	fmt.Printf("Emails (T)   : %T \n", emails)
	var startTime time.Time
	fmt.Printf("Start (#v)   : %#v \n", startTime)
	fmt.Printf("Start (+v)   : %+v \n", startTime)
	fmt.Printf("Start (T)    : %T \n", startTime)
}