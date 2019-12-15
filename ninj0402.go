// Using a COMPOSITE LITERAL:
// 	create a SLICE of TYPE int
// 	assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// 	print out the TYPE of the slice

// solution: https://play.golang.org/p/sAQeFB7DIs

package main

import (
	"fmt"
)

func main() {

	//slice idiom
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for k, v := range a {
		fmt.Printf("a at %v is %v\n", k, v)
	}

	fmt.Printf("The array type of a %T\n", a)
}
