package main

import (
	"fmt"
)

func main() {
	fmt.Println("Reflection is never clear.")

	x := "bob"
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sOb := []byte(x)
	fmt.Println(sOb)
	fmt.Printf("%T\n", sOb)
	for i := 0; i < len(sOb); i++ {
		fmt.Printf("For %q, the Binary value is %b, the Octal value is %o, the Decimal value is %d, the Hex value is %U \n", sOb[i], sOb[i], sOb[i], sOb[i], sOb[i])
	}
}

