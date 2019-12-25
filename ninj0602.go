//create a func with the identifier foo that
//takes in a variadic parameter of type int
//pass in a value of type []int into your func (unfurl the []int)
//returns the sum of all values of type int passed in

//create a func with the identifier bar that
//takes in a parameter of type []int
//returns the sum of all values of type int passed in

//code: https://play.golang.org/p/B0yRxtBQPD

package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 6}
	zad := foo(a...)
	fmt.Println("zad = ", zad)
	zed := bar(a)
	fmt.Println("zed = ", zed)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total = total + v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total = total + v
	}
	return total
}