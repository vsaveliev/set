package set

import (
	"fmt"
	"io"
	"sync"
)

// Print writes all generated sets from sequence of numbers in io.Writer
func Print(w io.Writer, sequence ...int) error {
	err := validateSequence(sequence)
	if err != nil {
		return fmt.Errorf("Cannot generate and print sets: %s", err)
	}

	max := sequence[0]
	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		set := make([]int, 0, len(sequence))
		set = append(set, i)

		wg.Add(1)
		go printReq(sequence[1:], set, w, &wg)
	}

	wg.Wait()

	return nil
}

func validateSequence(sequence []int) error {
	if len(sequence) < 2 {
		return fmt.Errorf("sequence size must be greater than 2")
	}

	for _, i := range sequence {
		if i <= 0 {
			return fmt.Errorf("sequence must have only positive numbers (>0)")
		}

		// TODO: limit sequence numbers' value
	}

	// TODO: limit sequence size

	return nil
}

func printReq(sequence, set []int, w io.Writer, wg *sync.WaitGroup) {
	defer wg.Done()

	max := sequence[0]
	for i := 0; i < max; i++ {
		if len(sequence) == 1 {
			fmt.Fprintln(w, append(set, i))
			continue
		}

		newSet := make([]int, len(set))
		copy(newSet, set)
		newSet = append(newSet, i)

		wg.Add(1)
		go printReq(sequence[1:], newSet, w, wg)
	}
}
