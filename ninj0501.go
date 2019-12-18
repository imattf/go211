//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:

//    first name
//    last name
//   favorite ice cream flavors

//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
//solution: https://play.golang.org/p/ouRHmH_POg

package main

import (
	"fmt"
)

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	p1 := person{
		firstName:   "bill",
		lastName:    "jones",
		favIceCream: []string{"oak", "pine"},
	}

	p2 := person{
		firstName:   "sam",
		lastName:    "jones",
		favIceCream: []string{"pine", "oaky"},
	}

	fmt.Println("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite!")

	fmt.Printf("%v's favortite ice-creams are:\n", p1.firstName)
	for k, v := range p1.favIceCream {
		fmt.Println(" ", k+1, v)
	}

	fmt.Printf("%v's favortite ice-creams are:\n", p2.firstName)
	for k, v := range p2.favIceCream {
		fmt.Println(" ", k+1, v)
	}
}
