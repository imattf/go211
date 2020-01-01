// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them

// solution: https://play.golang.org/p/096Lk1BR7o

package main

import (
	"fmt"
)

func main() {
	fmt.Println("main...")

	c := make(chan int)

	go put(c) //puts onto the channel

	pull(c) //pulls from the channel

	fmt.Println("...exiting main")
}

// put into the channel
func put(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Println("the value put in the channel:", i)
	}
	close(c)
}

// pull from the channel
func pull(c <-chan int) {
	for v := range c {
		fmt.Println("the value pulled from the channel:", v)
	}
}
