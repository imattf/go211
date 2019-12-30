//Starting with this code, marshal the []user to JSON.
//  There is a little bit of a curve ball in this hands-on exercise -
//  ---  remember to ask yourself what you need to do to EXPORT a value from a package.
// solution: https://play.golang.org/p/8BK6PXj3aG

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	u3 := user{"M", 54}
	users := []user{u1, u2, u3}

	fmt.Println(users)

	//  marshal the []user to JSON
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("json:", string(b))
	//os.Stdout.Write(b)
}
