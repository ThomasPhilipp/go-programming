package main

import (
	"fmt"
	"log"
	"os"
)

// Source: https://github.com/blackhat-go/bhg/blob/master/ch-2/io-example/main.go

// FooReader defines an io.Reader
type FooReader struct{}

// Read reads data from stdin.
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter defines an io.Writer
type FooWriter struct{}

// Write writes data to stdout.
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

// Reinvent Go's os.Stdin and os.Stdout types
func main() {

	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// Write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)

	// the previous code sample can be also replaced through -->
	//if _, err := io.Copy(&writer, &reader); err != nil {
	//	log.Fatalln("Unable to read/write data")
	//}
}