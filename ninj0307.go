// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
// solution: https://play.golang.org/p/IDnrJpE7vT

package main

import (
	"fmt"
)

func main() {

	i := 42
	if i == 40 {
		fmt.Println("is 40")
	} else if i == 41 {
		fmt.Println("is 41")
	} else {
		fmt.Println("is 42")
	}
}
