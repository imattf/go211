//various forms on the error format options

package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err01 happened", err)
		//log.Println("err02 happened", err)
		//log.Fatalln("err03 happened", err)
		//log.Panic("err04 happened ",err)
		//log.Panicln("err05 happened ", err)
		panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
