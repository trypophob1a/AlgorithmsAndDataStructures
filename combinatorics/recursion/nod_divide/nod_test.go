package noddivide

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNodeBySubtraction(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 4, b: 2, expected: 3},
		{a: 32, b: 15, expected: 1},
		{a: 16, b: 32, expected: 16},
		{a: 12, b: 5, expected: 1},
		{a: 3, b: 15, expected: 3},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(strconv.Itoa(tc.a), func(t *testing.T) {
			subtraction := getNodBySubtraction(tc.a, tc.b)
			mod := getNodByMod(tc.a, tc.b)
			require.Equal(t, tc.expected, subtraction)
			require.Equal(t, tc.expected, mod)
		})
	}
}
