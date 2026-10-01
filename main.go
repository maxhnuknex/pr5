package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Program supports Windows")
	case "linux":
		fmt.Println("Program supports Linux")
	default:
		fmt.Println("Unsupported operating system")
	}
}
