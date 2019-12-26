// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
// code: https://play.golang.org/p/c2HwqVE1Rd

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey...")

	a := func() func() {

		return func() {
			fmt.Println("oi")
		}

	}
	a()()

	b := foo()
	fmt.Println(b())

}

func foo() func() int {

	return func() int {
		return 42
	}
}
