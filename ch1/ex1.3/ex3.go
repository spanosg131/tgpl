// Package ex3 for benchmarking os.Args loops
// Exercise 1.3
package ex3

import (
	"os"
	"strings"
)

func argLoop() string {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s = sep + arg
	}
	return s
}

func joinLoop() string {
	return strings.Join(os.Args[1:], " ")
}
