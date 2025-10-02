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

var (
	flagFile = flag.String(
		"file",
		"",
		"File to open")
	flagUrl = flag.String(
		"url",
		"",
		"URL to load")
)

func main() {
	flag.Parse()

	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	switch {
	case *flagFile != "":
		fd, err := os.Open(*flagFile)
		if err != nil {
			fmt.Fprintln(output, "Error opening file: ", err)
			os.Exit(1)
		}
		defer fd.Close()

		input = fd

	case *flagUrl != "":
		resp, err := http.Get(*flagUrl)
		if err != nil {
			fmt.Fprintln(output, "Error loading url: ", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		input = resp.Body
	}

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
