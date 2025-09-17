package main

import (
	"fmt"
	"net"
	"sync"
)

// See also https://github.com/blackhat-go/bhg/tree/master/ch-2

// A TCP scanner which support concurrent calls and scale better. Putting
// "go" in front of the function is not enough and the app would immediately
// finish as it is not waiting to complete its single tasks. Therefore, we
// have to use the "WaitGroup".
func main()  {
	// Thread-safe way to control concurrency
	var wg sync.WaitGroup

	for i:=1; i<=65535; i++ {
		// Increase counter by the provided number
		wg.Add(1)

		go func(j int) {
			// Decrease counter by one whenever one unit of work has been performed
			defer wg.Done()

			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("Port %d close\n", j)
				return
			}
			defer conn.Close()
			fmt.Printf("Port %d open\n", j)
		}(i)
	}

	// Block the execution until the counter reaches zero ->  here the main
	// goroutine waits for all connections to finish
	wg.Wait()
}
