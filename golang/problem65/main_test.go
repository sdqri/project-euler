package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNthElement(t *testing.T) {
	elements := []int64{2, 1, 2, 1, 1, 4, 1, 1, 6, 1, 1, 8}

	for i, e := range elements {
		t.Run(fmt.Sprintf("%vth element of e", i), func(t *testing.T) {
			assert.Equal(t, e, GetNthElement(int64(i)))
		})
	}
}
