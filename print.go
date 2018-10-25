package set

import (
	"fmt"
	"io"
)

// Print writes all generated sets from sequence of numbers in io.Writer.
// The sets can be order randomly
func Print(w io.Writer, sequence ...int) error {
	results, err := Generate(sequence...)
	if err != nil {
		return err
	}

	for result := range results {
		fmt.Fprintln(w, result)
	}

	return nil
}
