package set

import (
	"fmt"
	"io"
	"sync"
)

// Print writes all generated sets from sequence of numbers in io.Writer
func Print(w io.Writer, sequence ...int) {
	max := sequence[0]
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		set := make([]int, 0, len(sequence))
		set = append(set, i)

		wg.Add(1)
		go printReq(sequence[1:], set, w, &wg)
	}

	wg.Wait()
}

func printReq(sequence, set []int, w io.Writer, wg *sync.WaitGroup) {
	defer wg.Done()

	max := sequence[0]
	for i := 0; i < max; i++ {
		newSet := make([]int, len(set))
		copy(newSet, set)
		newSet = append(newSet, i)

		if len(sequence) == 1 {
			fmt.Fprintln(w, newSet)
			continue
		}

		wg.Add(1)
		go printReq(sequence[1:], newSet, w, wg)
	}
}
