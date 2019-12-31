// Create a program that prints out your OS and ARCH. Use the following commands to run it
//     go run
//     go build
//     go install
// code: https://github.com/GoesToEleven/go-programming

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("in main...")
	fmt.Println("...CPUs:", runtime.NumCPU())
	fmt.Println("...the arch is:", runtime.GOARCH)
	fmt.Println("...the OS is:", runtime.GOOS)
}
