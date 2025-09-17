package main

import (
	"bufio"
	"log"
	"net"
)

// Note: net.Conn is a Reader and a Writer
func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()

	// Alternative approach: simply use the io.Copy function! See also IO.
}

// Echo server simply replies to all socket-request using the net.Conn type,
// which implements the Reader and Writer interfaces, as it is used for
// bidirectional communications: send (write) and receive (read) data.
//
// This example uses the bufio package, which creates a buffered I/O
// mechanism which simplify its logic.
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