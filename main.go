package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Program supports Windows")
	default:
		fmt.Println("Unsupported operating system")
	}
}
