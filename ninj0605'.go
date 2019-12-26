// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
// code: https://play.golang.org/p/NGGikHNvHv

package main

import (
	"fmt"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
	width  float64
}

func (c circle) area() float64 {
	return (3.14 * (c.radius * c.radius))
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case circle:
		totArea := 3.14 * s.(circle).radius * s.(circle).radius
		fmt.Println("i'm a circul", totArea)
	case square:
		totArea := s.(square).length * s.(square).width
		fmt.Println("i'm a skare", totArea)
	}
}

func main() {
	circle1 := circle{
		radius: 2,
	}

	square1 := square{
		length: 2,
		width:  2,
	}

	fmt.Println(circle1.area())
	fmt.Println(square1.area())

	info(circle1)
	info(square1)

}
