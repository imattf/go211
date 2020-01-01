// get this code working:

// using func literal, aka, anonymous self-executing func
//   solution: https://play.golang.org/p/SHr3lpX4so
// a buffered channel
//   solution: https://play.golang.org/p/Y0Hx6IZc3U

package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	//go func() {
	//	c <- 42
	//}()
	
	//or with...
	c <- 42

	fmt.Println(<-c)
}