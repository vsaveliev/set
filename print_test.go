package set

import (
	"io/ioutil"
	"testing"
)

// TestPrintEmptySequence checks that we cannot provide empty sequence to Print func
func TestPrintEmptySequence(t *testing.T) {
	w := ioutil.Discard
	err := Print(w)
	if err == nil {
		t.Errorf("Print must return the error if the sequence is empty")
	}
}

// TestPrintNegativeSequence checks that we cannot provide sequence with negative numbers to Print func
func TestPrintNegativeSequence(t *testing.T) {
	w := ioutil.Discard
	err := Print(w, 1, -2, 3)
	if err == nil {
		t.Errorf("Print must return the error if the sequence has negative numbers")
	}
}
