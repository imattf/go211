package main

import (
	"fmt"
)

func main() {
	fmt.Println("The bigger the interface, the weaker the abstraction.")
	a := 0
	b := 0

	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)

}
