// Ninja Exercise 01-05
package main

import (
	"fmt"
)

type mytype int

var x mytype
var y int

func main() {

	fmt.Println("x is of value", x)
	fmt.Printf("x is of type %T\n", x)
	x = 42
	fmt.Println("x is of value", x)

	y = int(x)
	fmt.Println("y is of value", y)
	fmt.Printf("y is of type %T\n", y)

}
