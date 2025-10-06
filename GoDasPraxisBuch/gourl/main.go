package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Example for loading a URL and store it to a file
//
// $ gourl -o <file.html> <url>
// $ gourl -h <url>

var (
	flagHeader = flag.Bool("h", false, "Print HTTP headers")
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

	if !validateUrl(url) {
		fmt.Printf("Error validating URL: %s\n", url)
		os.Exit(1)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error bei loading %s\n%#v", url, err)
	}
	defer resp.Body.Close()

	var w io.Writer
	w = os.Stdout

	if *flagHeader {
		for k, v := range resp.Header {
			fmt.Fprintf(w, "%s :\n", k)
			for i, l := range v {
				fmt.Fprintf(w, "	%03d: %s \n", i+1, l)
			}
		}
		os.Exit(0)
	}

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

func validateUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err !=  nil {
		return false
	}
	return true
}