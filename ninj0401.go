// Using a COMPOSITE LITERAL:
// 	create an ARRAY which holds 5 VALUES of TYPE int
// 	assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// 	print out the TYPE of the array

// solution: https://play.golang.org/p/tD0SzV3hdf

package main

import (
	"fmt"
)

func main() {
	//array idiom
	// var a [5]int
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3
	// a[3] = 4
	// a[4] = 5
	// fmt.Println(a)

	//slice idiom
	a := [5]int{1, 2, 3, 4, 5}

	for k, v := range a {
		fmt.Printf("a at %v is %v\n", k, v)
	}

	fmt.Printf("The array type of a %T\n", a)
}
