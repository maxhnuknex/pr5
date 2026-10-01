package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	default:
		fmt.Println("Unsupported operating system")
	}
}
