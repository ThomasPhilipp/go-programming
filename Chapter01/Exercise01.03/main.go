package main

import (
	"fmt"
	"time"
)

var (
	Debug bool = false
	LogLevel string = "info"
	statUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, statUpTime)
}
