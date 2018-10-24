package set_test

import (
	"os"

	"github.com/vsaveliev/set"
)

// ExamplPrint demonstrates work of Print
func ExamplePrint() {
	w := os.Stdout
	set.Print(w, 1, 1)
	// Output:
	// [0 0]
}
