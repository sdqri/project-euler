package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSquareRootPeriodLen(t *testing.T) {
	cases := []struct {
		num    int
		period int
	}{
		{2, 1},
		{3, 2},
		{5, 1},
		{6, 2},
		{7, 4},
		{8, 2},
		{10, 1},
		{11, 2},
		{12, 2},
		{13, 5},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("sqrt(%d)", tc.num), func(t *testing.T) {
			assert.Equal(t, GetSquareRootPeriodLen(float64(tc.num)), tc.period)
		})
	}
}
