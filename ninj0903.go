// Hands-on Exercise #3; Ninja Level 9
// Using goroutines, create an incrementer program
// have a variable to hold the incrementer value
// launch a bunch of goroutines
// each goroutine should
// read the incrementer value
// store it in a new variable
// yield the processor with runtime.Gosched()
// increment the new variable
// write the value in the new variable back to the incrementer variable
// use waitgroups to wait for all of your goroutines to finish
// the above will create a race condition.
// Prove that it is a race condition by using the -race flag
// if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
// code: https://github.com/GoesToEleven/go-programming

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var increment int

func main() {
	fmt.Println("in main...")
	fmt.Println("...CPUs:", runtime.NumCPU())
	fmt.Println("...and count is:", increment)

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go foo()
	}
	fmt.Println("...and now count is:", increment)
	wg.Wait()

	fmt.Println("...and finally count is:", increment)
}

func foo() {

	v := increment
	runtime.Gosched()
	v++
	increment = v
	fmt.Println("...and now count is:", increment)
	wg.Done()

}
