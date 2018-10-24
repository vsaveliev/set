package set

import (
	"io/ioutil"
	"testing"
)

func BenchmarkPrint3(b *testing.B) {
	benchmarkPrint(3, b)
}

func BenchmarkPrint5(b *testing.B) {
	benchmarkPrint(5, b)
}

func BenchmarkPrint7(b *testing.B) {
	benchmarkPrint(7, b)
}

func benchmarkPrint(seqSize int, b *testing.B) {
	seq := make([]int, seqSize)
	for i := 0; i < seqSize; i++ {
		seq[i] = 10
	}

	w := ioutil.Discard

	for i := 0; i < b.N; i++ {
		Print(w, seq...)
	}
}
