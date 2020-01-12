package main

import (
	"fmt"

	"github.com/imattf/go211/ninj1300/02/quote"
	"github.com/imattf/go211/ninj1300/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
