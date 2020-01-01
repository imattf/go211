// Starting with this code, pull the values off the channel using a select statement
// solution: https://play.golang.org/p/FulKBY5JNj

package main

import (
	"fmt"
)

func main() {

	bq := make(chan int) //bi-directional quit
	cr := gen(bq)        //receive chan

	receive(cr, bq)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
