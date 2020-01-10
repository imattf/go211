// Start with this code. Create a custom error message using “fmt.Errorf”.
// solution:
//     https://play.golang.org/p/HugU4HJEEO
//     https://play.golang.org/p/NII-lmGasj
//     https://play.golang.org/p/Vo5kIoR-sG

package main

import (
	"encoding/json"
	"errors"
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

	bs, err := toJSON(p1)
	if err != nil {
		log.Panicln(err)

	}

	fmt.Println(string(bs), err)

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		//return bs, fmt.Errorf("toJSON error: %v", err)
		return bs, errors.New(fmt.Sprintf("toJSON error %v", err))
	}
	return bs, nil
}
