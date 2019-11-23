// Ninja Exercise 01-03
package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprintf("x is: %d and y is: %s and z is: %t", x, y, z)
	fmt.Println(s)

}
