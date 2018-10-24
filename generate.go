package set

import (
	"fmt"
	"sync"
)

// Generate returns channel with generated sets from sequence
func Generate(sequence ...int) (<-chan []int, error) {
	err := validateSequence(sequence)
	if err != nil {
		return nil, fmt.Errorf("Cannot generate and print sets: %s", err)
	}

	results := make(chan []int)

	go func() {
		defer close(results)

		max := sequence[0]
		var wg sync.WaitGroup
		for i := 0; i < max; i++ {
			set := make([]int, 0, len(sequence))
			set = append(set, i)

			wg.Add(1)
			go generateReq(sequence[1:], set, results, &wg)
		}

		wg.Wait()
	}()

	return results, nil
}

func generateReq(sequence, set []int, results chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

	max := sequence[0]
	for i := 0; i < max; i++ {
		newSet := make([]int, len(set))
		copy(newSet, set)
		newSet = append(newSet, i)

		if len(sequence) == 1 {
			results <- newSet
			continue
		}

		wg.Add(1)
		go generateReq(sequence[1:], newSet, results, wg)
	}
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
