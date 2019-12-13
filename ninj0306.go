// Create a program that shows the “if statement” in action.
// solution: https://play.golang.org/p/DpZ_FLfn5s

package main

import (
	"fmt"
)

func main() {

	for i := 10; i <= 100; i++ {
		fmt.Printf("%v modulo 4 is %v\t\n", i, i%4)

	}

}
