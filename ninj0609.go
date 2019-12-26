// A “callback” is when we pass a func into a func as an argument. For this exercise,
//     pass a func into a func as an argument
// code: https://play.golang.org/p/0yGYPKh1y7

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey...")
	bar(foo)
}

func foo() {
	fmt.Println("...in foo")
}

func bar(f func()) {
	fmt.Println("...in bar")
	f()
}
