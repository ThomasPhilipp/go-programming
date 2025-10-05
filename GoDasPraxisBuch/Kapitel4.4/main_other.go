// +build darwin linux

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("ifconfig", "-a").Output()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("IP address on OS X", out)
}