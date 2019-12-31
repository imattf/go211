// Fix the race condition you created in the previous exercise by using a mutex
//     it makes sense to remove runtime.Gosched()
// code: https://github.com/GoesToEleven/go-programming

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var increment int

func main() {
	fmt.Println("in main...")
	fmt.Println("...CPUs:", runtime.NumCPU())

	fmt.Println("...count starts as:", increment)

	wg.Add(100)

	for i := 0; i < 100; i++ {

		go foo()
		// go func() {
		// 	mu.Lock()
		// 	v := increment
		// 	//runtime.Gosched()
		// 	v++
		// 	increment = v
		// 	fmt.Println("...and now count is:", increment)
		// 	mu.Unlock()
		// 	wg.Done()
		// }()
	}
	wg.Wait()

	fmt.Println("...and finally count is:", increment)
}

func foo() {

	mu.Lock()
	v := increment
	runtime.Gosched()
	v++
	increment = v
	fmt.Println("...and now count is:", increment)
	mu.Unlock()
	wg.Done()

}
