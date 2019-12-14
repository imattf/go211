//Create a program that uses a switch statement with no switch expression specified.
//solution: https://play.golang.org/p/YpPgkWGqKY

package main

import (
	"fmt"
)

func main() {

	i, j := 42, 42
	switch {
	case i < j:
		fmt.Println("<")
	case i > j:
		fmt.Println(">")
	case i == j:
		fmt.Println("=")
	}
}
