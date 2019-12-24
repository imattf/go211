// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(foo())
	x, y := bar()
	fmt.Println(y, x)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "hi"
}
