// Exercise 1.2
// Display each os.Arg argument one per line with their respective index number
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for index, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(index, arg)
	}
}
