package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "joescatcam.website:80")
	if err != nil {
		log.Fatalln("Unable to connect to host")
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy src output to dst (aka data from the inbound connection is copied to dst)
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy dst output back to src (aka data read from dst is written back to client's connection)
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

// A proxy to circumvent restrictive egress controls, e.g. through port
// forwarding, where all traffic is redirected:
// * Access to joescatcam.website prohibited
// * Access to joesproxy.com allowed
// * Forward traffic at joesproxy.com to joescatcam.website
func main() {
	// Listen on local port 80, e.g. joesproxy.com
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}