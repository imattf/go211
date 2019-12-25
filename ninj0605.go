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

/// shape interface goes here
type shape interface {
	area() float64
}

/// info method goes here
func info(s shape) {
	fmt.Println(s.area())
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
