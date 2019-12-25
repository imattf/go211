// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
// code: https://play.golang.org/p/XluEuUD0Nw

package main

func main() {
	defer foo()
	bar()
}

func foo() {
	println("foo")
}

func bar() {
	println("bar")
}
