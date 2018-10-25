package set

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// TestGenerate checks generated sets from different sequences
func TestGenerate(t *testing.T) {
	// TODO: move fixtures in separate file (yml / json / ...)
	var fixtures = []struct {
		input    []int
		expected []string
	}{
		{
			[]int{1, 2, 3},
			[]string{"[0 0 0]", "[0 0 1]", "[0 0 2]", "[0 1 0]", "[0 1 1]", "[0 1 2]"},
		},
		{
			[]int{1, 2},
			[]string{"[0 0]", "[0 1]"},
		},
		{
			[]int{2, 2},
			[]string{"[0 0]", "[0 1]", "[1 0]", "[1 1]"},
		},
	}

	for _, fixture := range fixtures {
		testGeneration(t, fixture.input, fixture.expected)
	}
}

func testGeneration(t *testing.T, sequence []int, expectedResults []string) {
	resultsChan, err := Generate(sequence...)
	if err != nil {
		t.Errorf("Print must process request without any errors, but: %s", err)
	}

	// it's easy to check slice of string
	realResults := make([]string, 0, len(expectedResults))
	for set := range resultsChan {
		realResults = append(realResults, fmt.Sprint(set))
	}
	sort.Strings(realResults)

	if !reflect.DeepEqual(expectedResults, realResults) {
		t.Errorf("Result is unexpected: %v, expected %v", realResults, expectedResults)
	}
}
