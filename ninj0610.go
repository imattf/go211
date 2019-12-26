// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable:
// code: https://play.golang.org/p/a56uWtoFSL

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey...")
	a := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := foo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
