// Fix the race condition you created in exercise #4 by using package atomic
// code: https://github.com/GoesToEleven/go-programming

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var increment int64

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

	atomic.AddInt64(&increment, 1)
	runtime.Gosched()
	fmt.Println("...and now count is:", atomic.LoadInt64(&increment))
	wg.Done()

}
