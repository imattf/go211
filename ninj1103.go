// Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
// solution:
//     https://play.golang.org/p/ixeowY2fd2
//     assertion
//     https://play.golang.org/p/pbl2kCYsM0
//     conversion
//     https://play.golang.org/p/1ldiBdkdzA

package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("error happened: %v", ce.info)
}

func main() {
	ce := customErr{
		info: "i am ce1",
	}

	foo(ce)
}

func foo(e error) {
	fmt.Println("in foo", e)
}
