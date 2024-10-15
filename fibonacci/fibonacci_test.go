package fibonacci

import (
	"testing"
)

func TestFibN(t *testing.T) {
	result := fibN(6)

	if result != 8 {
		t.Error("Expected value for fib 6 is not equal to result: ", result)
	}
}
