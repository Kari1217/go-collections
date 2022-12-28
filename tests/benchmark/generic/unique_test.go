package generic

import (
	"testing"

	. "github.com/thefuga/go-collections"
	"github.com/thefuga/go-collections/tests/benchmark"
)

func BenchmarkUnique(b *testing.B) {
	slice := benchmark.BuildIntSlice()

	for n := 0; n < b.N; n++ {
		Unique(slice)
	}
}
