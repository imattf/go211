// Create a value and assign it to a variable.
// Print the address of that value.
// code: https://play.golang.org/p/57kW8xd0qT

package main

import "fmt"

func main() {
	x := 42
	fmt.Println("the address of x is", &x)
}
