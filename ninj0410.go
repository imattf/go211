package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}

	m["jaws_jaws"] = []string{"Being not so Evil", "Ice Cream", "Sunsets"}

	for k, v := range m {
		fmt.Printf("%v favorite things are:\n", k)
		for k1, v1 := range v {
			fmt.Printf("   %v - %v\n", k1, v1)
		}
	}

	fmt.Println("Hello, deletion...")

	if v, ok := m["jaws_jaws"]; ok {
		fmt.Println("found", v)
		delete(m, "jaws_jaws")
	} else {
		fmt.Println("not found")
	}

	for k, v := range m {
		fmt.Printf("%v favorite things are:\n", k)
		for k1, v1 := range v {
			fmt.Printf("   %v - %v\n", k1, v1)
		}
	}
}
