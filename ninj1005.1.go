package main

import (
	"fmt"
)

func main() {

	var ans1, ans2, ans3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		fmt.Println("shit", err)
		panic(err)
	}
	fmt.Print("Drink: ")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		fmt.Println("shit", err)
		panic(err)
	}
	fmt.Print("Skills: ")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		fmt.Println("shit", err)
		panic(err)
	}

	//final
	fmt.Println("workd", ans1, ans2, ans3)

}
