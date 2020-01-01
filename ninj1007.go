// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them
// solutions:
// https://play.golang.org/p/R-zqsKS03P
// https://play.golang.org/p/quWnlwzs2z

package main

import (
	"fmt"
)

func main() {
	fmt.Println("main...")

	c := make(chan int, 10)

	//puts onto the channel
	for i := 0; i < 10; i++ {
		go put(c)
	}

	//pulls from the channel
	for i := 0; i < 100; i++ {
		fmt.Printf("starting with %v the value pulled from the channel:%v\n", i, <-c)
	}

	fmt.Println("...exit main")
}

// put into the channel
func put(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("the value put in the channel:", i)
	}
}
