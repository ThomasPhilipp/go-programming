package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	// Create a buffer to store received data
	b := make([]byte, 4096)

	for {
		// Receive data and sotre it into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}

		if err != nil {
			log.Println("Unexpected error")
		}

		log.Printf("Received %d bytes: %s\n", size, string(b))

		// Send data
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

// Echo server simply replies to all socket-request using the net.Conn type,
// which implements the Reader and Writer interfaces, as it is used for
// bidirectional communications: send (write) and receive (read) data.
func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:8080")

	for {
		// Wait for client connections
		conn, err := listener.Accept() // it blocks until a client connects and returns a Conn instance
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection
		go echo(conn)
	}
}