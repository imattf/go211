//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
//solution: https://play.golang.org/p/h-8FnjpzWB

package main

import (
	"fmt"
)

func main() {

	var favSpot string = "declining"

	switch favSpot {
	case "reclining":
		fmt.Println("...in a chair")
	case "inclining":
		fmt.Println("...on a hill")
	case "declining":
		fmt.Println("...at a meal")
	}
}
