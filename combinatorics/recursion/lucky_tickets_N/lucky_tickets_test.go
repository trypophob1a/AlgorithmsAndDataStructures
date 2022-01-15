package luckyticketsN

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLuckyTicketsCount(t *testing.T) {
	tests := []struct {
		input    uint
		expected uint
	}{
		{input: 3, expected: 55252},
		{input: 1, expected: 10},
		{input: 2, expected: 670},
		{input: 4, expected: 4816030},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(strconv.Itoa(int(tc.input)), func(t *testing.T) {
			result := luckyTicketsCount(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
