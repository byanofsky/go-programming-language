package echo

import (
	"testing"
)

var args = []string{"1", "2", "3", "4", "5"}

func BenchmarkInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoInefficient(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoJoin(args)
	}
}
