// Start with this code. Instead of using the blank identifier, make sure the code is checking and handling the error.
// solution:
//     https://play.golang.org/p/tn8oiuL1Yn

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		//fmt.Println(err)
		//or...
		log.Fatalln("JSON did not marshal - here's the error: ", err)
	}
	fmt.Println(string(bs))

}
