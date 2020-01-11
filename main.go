package main

import (
	"fmt"

	"./dog"
)

type human struct {
	name string
	age  int
}

func main() {
	dad := human{
		name: "Albert",
		age:  67,
	}
	dogYears := dog.Years(dad.age)
	fmt.Println(dogYears)
}
