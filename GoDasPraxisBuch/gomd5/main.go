package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// A Go-based command implementation to create a MD5 hash from the standard
// input.
// $ gomd5 <file>
// $ gomd5 <url>
// $ gomd5 -h

// Define command-line flags for input sources
var (
	// -file flag: specifies a local file to open
	flagFile = flag.String(
		"file",
		"",
		"File to open")
	// -url flag: specifies a URL to fetch data from
	flagUrl = flag.String(
		"url",
		"",
		"URL to load")
)

func main() {
	// Parse the command-line flags provided by the user
	flag.Parse()

	// Set default input and output streams
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	// Determine the input source based on the flags
	switch {
	case *flagFile != "":
		// If -file flag is provided, try to open the specified file
		fd, err := os.Open(*flagFile)
		if err != nil {
			fmt.Fprintln(output, "Error opening file: ", err)
			os.Exit(1)
		}
		// Ensure the file is closed when done
		defer fd.Close()

		// Set input to the opened file
		input = fd

	case *flagUrl != "":
		// If -url flag is provided, try to fetch the content from the URL
		resp, err := http.Get(*flagUrl)
		if err != nil {
			fmt.Fprintln(output, "Error loading url: ", err)
			os.Exit(1)
		}
		// Ensure the response body is closed when done
		defer resp.Body.Close()

		// Set input to the response body
		input = resp.Body
	}

	// Compute and print the MD5 hash of the input data
	printMD5(input, os.Stdout)
}

func printMD5(i io.Reader, w io.Writer) {
	// Create a new MD5 hash object
	h := md5.New()

	// Copy data from standard input (os.Stdin) into the hash object
	// This reads all input and updates the hash state
	if _, err := io.Copy(h, i); err != nil {
		log.Fatal("Error during reading: ", err)
		return
	}

	// Compute the final MD5 checksum and print it as a hexadecimal string
	fmt.Fprintf(w, "%x", h.Sum(nil))
}
