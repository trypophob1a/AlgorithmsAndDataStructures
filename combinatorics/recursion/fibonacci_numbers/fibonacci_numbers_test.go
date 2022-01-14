package fibonaccinumbers

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNaiveAlgorithm(t *testing.T) {
	tests := []struct {
		input    uint
		expected uint
	}{
		{input: 4, expected: 3},
		{input: 7, expected: 13},
		{input: 11, expected: 89},
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 2, expected: 1},
		{input: 5, expected: 5},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(strconv.Itoa(int(tc.input)), func(t *testing.T) {
			result := naiveAlgorithm(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
