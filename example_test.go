package set_test

import (
	"fmt"
	"os"

	"github.com/vsaveliev/set"
)

// ExamplPrint demonstrates work of Print func
func ExamplePrint() {
	w := os.Stdout
	err := set.Print(w, 1, 1)
	if err != nil {
		fmt.Println("Cannot print sets:", err)
	}
	// Output:
	// [0 0]
}

// ExampleGenerate demonstrates work of Generate func
func ExampleGenerate() {
	results, err := set.Generate(1, 1)
	if err != nil {
		fmt.Println("Cannot generate sets:", err)
	}

	for result := range results {
		fmt.Println(result)
	}
	// Output:
	// [0 0]
}
