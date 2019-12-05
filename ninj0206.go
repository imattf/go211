package main

import (
	"fmt"
)

const (
	year4 = 2019 - iota
	year3
	year2
	year1
)

func main() {
	fmt.Println("The Bigger the interface, the weaker the abstraction.")

	fmt.Println(year4)
	fmt.Println(year3)
	fmt.Println(year2)
	fmt.Println(year1)
}

// Using iota, create 4 constants for the last 4 years. Print the constant values.
