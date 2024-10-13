package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenFactorialChains(t *testing.T) {
	testCases := []struct {
		x           int64
		chainLength int64
	}{
		{145, 1},
		{169, 3},
		{871, 2},
		{872, 2},
		{69, 5},
		{78, 4},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("x=%d", tc.x), func(t *testing.T) {
			actual := LenDigitFactorialChain(nil, tc.x)
			assert.Equal(t, tc.chainLength, actual)
		})
	}
}
