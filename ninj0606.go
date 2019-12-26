// Build and use an anonymous func
// code: https://play.golang.org/p/DQX3xEIcRe

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey...")
	func() {
		fmt.Println("...anonymously")
	}()
}
