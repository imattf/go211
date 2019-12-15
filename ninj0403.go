// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
//     [42 43 44 45 46]
//     [47 48 49 50 51]
//     [44 45 46 47 48]
//     [43 44 45 46 47]
// solution: https://play.golang.org/p/SGfiULXzAB

package main

import (
	"fmt"
)

func main() {

	//new slice idiom
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	// a1 := a[:5]
	// fmt.Println(a1)
	// a2 := a[5:]
	// fmt.Println(a2)
	// a3 := a[2:7]
	// fmt.Println(a3)
	// a4 := a[1:6]
	// fmt.Println(a4)

	//slice of slice idiom
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(a[:5])
	fmt.Println(a[5:])
	fmt.Println(a[2:7])
	fmt.Println(a[1:6])
}
