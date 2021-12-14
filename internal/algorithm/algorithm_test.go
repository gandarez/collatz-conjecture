package algorithm_test

import (
	"testing"

	"github.com/gandarez/collatz-conjecture/internal/algorithm"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := map[string]struct {
		Input    int
		Expected []int
	}{
		"number 7": {
			Input:    7,
			Expected: []int{22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1},
		},
		"number 1": {
			Input:    1,
			Expected: []int{4, 2, 1},
		},
		"number 2": {
			Input:    2,
			Expected: []int{1},
		},
		"number 4": {
			Input:    4,
			Expected: []int{2, 1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			results := algorithm.Calculate(test.Input)

			assert.Equal(t, test.Expected, results)
		})
	}
}
