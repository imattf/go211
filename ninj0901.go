// In addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

// code: https://github.com/GoesToEleven/go-programming

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go foo()

	wg.Add(1)
	go bar()

	wg.Wait()
}

func foo() {
	fmt.Println("...in foo")
	wg.Done()
}
func bar() {
	fmt.Println("...in bar")
	wg.Done()
}
