package palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		a        string
		expected bool
	}{
		{a: "Ага", expected: true},
		{a: "", expected: true},
		{a: "0000", expected: true},
		{a: "дохоД", expected: true},
		{a: "загадка", expected: false},
		{a: "1010", expected: false},
		{a: "2233", expected: false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.a, func(t *testing.T) {
			result := isPalindrome(tc.a)
			require.Equal(t, tc.expected, result)
		})
	}
}
