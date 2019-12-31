// The best way to learn is to teach. For this hands-on exercise,
//     choose one of the above exercises, or use the recursion example of factorial
//     download, install, and get it running
//     https://obsproject.com/
//     record a video of YOU teaching the topic
//     upload the video to youtube
// 	share the video on twitter and tag me in it
// 	( https://twitter.com/Todd_McLeod ) so that I can see it!

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey...")
	a := foo()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := foo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
