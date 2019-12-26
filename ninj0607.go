// Assign a func to a variable, then call that func
// code: https://play.golang.org/p/_Qu7ZAyFDH

package main

import (
	"fmt"
)

func main() {

	x := func() {
		fmt.Println("bark")
	}
	x()
	dog()
	c := dog
	c()

}

func dog() {
	println("bark, bark")
}
