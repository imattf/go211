// Ninja Exercise 01-04
package main

import (
	"fmt"
)

type mytype int

var x mytype

func main() {

        fmt.Println("x is of value",x)
	fmt.Printf("x is of type %T\n", x)
	x = 42
	fmt.Println("x is of value", x)

}
