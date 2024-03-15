package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsDigitCancelling(t *testing.T) {
	cases := map[string]struct {
		numerator   int
		denominator int
		expected    bool
	}{
		"49/98":  {numerator: 49, denominator: 98, expected: true},
		"98/48":  {numerator: 98, denominator: 48, expected: false},
		"26/65":  {numerator: 26, denominator: 65, expected: false},
		"499/98": {numerator: 499, denominator: 98, expected: false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual, err := IsDigitCancelling(tc.numerator, tc.denominator, false)
			if err != nil {
				t.Error(err)
			}
			require.Equal(t, tc.expected, actual)
		})
	}
}
