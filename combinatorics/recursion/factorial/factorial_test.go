package factorial

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		a        int
		expected int
	}{
		{a: 3, expected: 6},
		{a: 1, expected: 1},
		{a: 2, expected: 2},
		{a: 5, expected: 120},
		{a: 10, expected: 3628800},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(strconv.Itoa(tc.a), func(t *testing.T) {
			naiveFactorial := factorial(tc.a)
			tailRecursion := factorialWithTailRecursion(tc.a, 1)
			require.Equal(t, tc.expected, naiveFactorial)
			require.Equal(t, tc.expected, tailRecursion)
		})
	}
}
