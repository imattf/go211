// This exercise will reinforce our understanding of method sets:

//     create a type person struct
//     attach a method speak to type person using a pointer receiver
//     *person
//     create a type human interface
//     to implicitly implement the interface, a human must have the speak method
//     create func “saySomething”
//     have it take in a human as a parameter
//     have it call the speak method
//     show the following in your code
//     you CAN pass a value of type *person into saySomething
//     you CANNOT pass a value of type person into saySomething
//     here is a hint if you need some help
//     https://play.golang.org/p/FAwcQbNtMG

// Receivers       Values
// -----------------------------------------------
// (t T)           T and *T
// (t *T)          *T

// code: https://github.com/GoesToEleven/go-programming

package main

import "fmt"

type person struct {
	name string
	age  int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.name)
	fmt.Println("...i am", p.age, "years old")
}

//     create func “saySomething”
//     have it take in a human as a parameter
//     have it call the speak method
func saySomething(h human) {
	h.speak()
}

func main() {
	fmt.Println("Main...")
	a := person{"Sam", 32}
	//a.speak()

	//this will work...
	saySomething(&a)

	//this will not work...
	//saySomething(a)
}
