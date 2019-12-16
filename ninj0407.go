// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
// 		"James", "Bond", "Shaken, not stirred"
// 		"Miss", "Moneypenny", "Helloooooo, James."
// Range over the records, then range over the data in each record.
// solution: https://play.golang.org/p/FMM4c2PhZg

package main

import (
	"fmt"
)

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xy := [][]string{x, y}
	fmt.Println(xy)

	for k, v := range xy {
		fmt.Println(k, v)
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
}
