package main

import (
	"fmt"
	"sync"
)

// See also https://github.com/blackhat-go/bhg/tree/master/ch-2

// The channel is used to receive work and the WaitGroup to track when a single
// work item has been completed.
func worker(ports chan int, wg *sync.WaitGroup) {
	// Continuously receive from the ports channel until channel is closed
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

// A TCP scanner which support concurrent calls and scale better. Putting
// "go" in front of the function is not enough and the app would immediately
// finish as it is not waiting to complete its single tasks. Therefore, we
// have to use the "WaitGroup". Additionally, we create a worker pool and in
// the main the workload is managed.
func main()  {
	// Buffered channel which stores items without waiting for a receiver to
	// read them. It can hold 100 items before the sender will block.
	ports := make(chan int, 100)

	// Thread-safe way to control concurrency
	var wg sync.WaitGroup

	// Create a resource pool of workers (100) which consume a channel
	for i:=0; i<cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i:=1; i<=65535; i++ {
		// Increase counter by the provided number
		wg.Add(1)

		// Send a port to the ports channel to the worker
		ports <- i
	}

	// Block the execution until the counter reaches zero ->  here the main
	// goroutine waits for all connections to finish
	wg.Wait()

	// After the work has been completed, close the channel
	close(ports)
}
