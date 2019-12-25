// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person

// code: https://play.golang.org/p/NnXyWdqbbw

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//my first attempt; didn't leverage the coolness of the struct type in the TYPE method
//func (p person) iAm(n, l string, years int) {
//	fmt.Println("Hello, my name is", n, l, "and I am", years,"years old.")
//}

func (p person) iAm() {
	fmt.Println("Hello, my name is", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	p1 :=
		person{
			first: "bob",
			last:  "mendez",
			age:   55,
		}

	p1.iAm()

}
