package set

import (
	"fmt"
	"io"
)

// Print writes all generated sets from number sequence in io.Writer
func Print(w io.Writer, sequence ...int) {
	max := sequence[0]

	for i := 0; i < max; i++ {
		set := make([]int, 0, len(sequence))
		set = append(set, i)

		printReq(sequence[1:], set, w)
	}
}

func printReq(sequence, set []int, w io.Writer) {
	max := sequence[0]

	for i := 0; i < max; i++ {
		newSet := make([]int, len(set))
		copy(newSet, set)
		newSet = append(newSet, i)

		if len(sequence) == 1 {
			fmt.Fprintln(w, newSet)
			continue
		}

		printReq(sequence[1:], newSet, w)
	}
}
