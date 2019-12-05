package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reflection is never clear.")

	x := 42
	fmt.Printf("For %v, the Binary value is %b, the Octal value is %o, the Decimal value is %d, the Hex value is %U \n", x, x, x, x, x)
	y := x << 1
	fmt.Printf("For %v, the Binary value is %b, the Octal value is %o, the Decimal value is %d, the Hex value is %U \n", y, y, y, y, y)

}

// Write a program that

//     assigns an int to a variable
//     prints that int in decimal, binary, and hex
//     shifts the bits of that int over 1 position to the left, and assigns that to a variable
//     prints that variable in decimal, binary, and hex
