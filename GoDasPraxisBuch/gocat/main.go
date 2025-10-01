package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// A Go-based "cat" command implementation, which reads files (input) and
//prints out to stdout (output). Use it by calling:
// $ gocat <file1> <file2>

func main() {
	// Check if no command-line arguments (other than the program name) are provided
	if len(os.Args) == 1 {
		fmt.Println("Please specify minimum one file")
		// Exit the program with a non-zero status to indicate an error
		os.Exit(1)
	}

	// Iterate over each file name provided as a command-line argument
	for _, fName := range os.Args[1:] {
		// Attempt to open the file
		f, err := os.Open(fName)
		if err != nil {
			log.Printf("Error opening file: %s\n", fName, err)
			// Attempt to close the file (though f may be nil if open failed)
			f.Close()
			// Skip to the next file
			continue
		}

		// Copy the contents of the file to standard output
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Printf("Error during printing %s\n%s", fName, err)
		}
		// Close the file after processing
		f.Close()
	}
}
