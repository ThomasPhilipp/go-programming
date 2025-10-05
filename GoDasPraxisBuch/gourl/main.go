package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Example for loading a URL and store it to a file
//
// $ gourl -o <file.html> <url>

var (
	flagOutput = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Please specify only 1 URL")
		os.Exit(1)
	}

	url := args[0]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error bei loading %s\n%#v", url, err)
	}
	defer resp.Body.Close()

	var w io.Writer
	w = os.Stdout

	if *flagOutput != "" {
		err := os.MkdirAll(filepath.Dir(*flagOutput), 0755)
		if err != nil {
			fmt.Printf("Error by creating folder(s): %v", err)
			os.Exit(1)
		}

		f, err := os.OpenFile(*flagOutput, os.O_RDWR | os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Error by file creating: %s\n%hv", *flagOutput, err)
			os.Exit(1)
		}
		defer f.Close()

		w = f
	}

	io.Copy(w, resp.Body)
}