package main

import (
	"fmt"
	"os"
)

var (
	home = os.Getenv("HOME")
	user = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

func main() {
	fmt.Println("HOME variable: ", home)
	fmt.Println("USER variable: ", user)
	fmt.Println("GOPATH variable: ", gopath)
}