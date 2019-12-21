// Create and use an anonymous struct.
// solution: https://play.golang.org/p/otBHFs-lPp

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, dog playground")

	dog := struct {
		breed    string
		sound    string
		legs     map[string]int
		markings []string
	}{
		breed: "poodle",
		sound: "bark",
		legs: map[string]int{
			"left-front":  15,
			"right-front": 15,
			"left-back":   15,
			"right-back":  15,
		},
		markings: []string{"red spots", "curly hair"},
	}

	fmt.Println(dog.breed)
	fmt.Println(dog.sound)
	fmt.Println(dog.legs)
	fmt.Println(dog.markings)

	for k, v := range dog.legs {
		fmt.Println(k, v)
	}

	for i, v := range dog.markings {
		fmt.Println(i, v)
	}
}

//notice we can re-use v in a for statement
