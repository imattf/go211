//We are going to learn about testing next.
//For this exercise, take a moment and see how much you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s article on testing.
//See if you can get a basic example of testing working.
// https://godoc.org/testing
// https://www.golang-book.com/books/intro/12
//good tutorial on testing: https://tutorialedge.net/golang/intro-testing-in-go/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
