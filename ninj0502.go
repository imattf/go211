// Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
// Access each value in the map. Print out the values, ranging over the slice.
// solution: https://play.golang.org/p/RvvLyAMvGo

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	favFlav   []string
}

func main() {

	p1 := person{
		firstName: "bill",
		lastName:  "jones",
		favFlav:   []string{"oak", "pine"},
	}

	p2 := person{
		firstName: "sam",
		lastName:  "sneed",
		favFlav:   []string{"pine", "oaky"},
	}

	fmt.Println("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite!")

	fmt.Printf("%v's favortite ice-creams are:\n", p1.firstName)
	for k, v := range p1.favFlav {
		fmt.Println(" ", k+1, v)
	}

	fmt.Printf("%v's favortite ice-creams are:\n", p2.firstName)
	for k, v := range p2.favFlav {
		fmt.Println(" ", k+1, v)
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Printf("%v %v's favorite ice cream flavors are:\n", v.firstName, v.lastName)
		for i, val := range v.favFlav {
			fmt.Println(" ", i+1, val)

		}

	}

}
