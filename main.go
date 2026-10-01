package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("Program supports Linux")
	default:
		fmt.Println("Unsupported operating system")
	}
}
