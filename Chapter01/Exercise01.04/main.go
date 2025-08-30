package main

import (
	"fmt"
	"time"
)

var (
	Debug bool // omit initial value
	LogLevel = "info" // omit data type
	statUpTime = time.Now() // omit data type
)

func main() {
	fmt.Println(Debug, LogLevel, statUpTime)
}
