package main

import (
	"fmt"
	"net"
)

// See also https://github.com/blackhat-go/bhg/tree/master/ch-2

// A simple TCP scanner which connects to each port in the loop separately. It
// does not support concurrent calls and does not scale.
func main()  {
	for i:=0; i<=10; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Port %d close", i)
		}
		conn.Close()
		fmt.Printf("Port %d open\n", i)
	}
}
