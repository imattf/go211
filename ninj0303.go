// Create a for loop using this syntax
//     for condition { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
)

func main() {
	i := 1961
	for i <= 2019 {
		fmt.Println(i)
		i++
	}
}
