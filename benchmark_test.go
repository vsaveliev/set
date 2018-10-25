package set

import (
	"io/ioutil"
	"testing"
)

// BenchmarkPrint3 tests sets generation 3*N
func BenchmarkPrint3(b *testing.B) {
	benchmarkPrint(3, b)
}

// BenchmarkPrint4 tests sets generation 4*N
func BenchmarkPrint4(b *testing.B) {
	benchmarkPrint(4, b)
}

// BenchmarkPrint5 tests sets generation 5*N
func BenchmarkPrint5(b *testing.B) {
	benchmarkPrint(5, b)
}

func benchmarkPrint(seqSize int, b *testing.B) {
	seq := generateSequence(seqSize)

	// we discard all output
	w := ioutil.Discard

	for i := 0; i < b.N; i++ {
		Print(w, seq...)
	}
}

// generateSequence returns the sequence for tests
func generateSequence(size int) []int {
	seq := make([]int, size)
	for i := 0; i < size; i++ {
		// TODO: generate number randomly
		seq[i] = 5
	}
	return seq
}
