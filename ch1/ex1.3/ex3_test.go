// exercise 3
// Benchmark os.Args output with ange and with strings.Join
package ex3

import (
	"testing"
)

func BenchmarkArgLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		argLoop()
	}
}

func BenchmarkJoinLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinLoop()
	}
}
