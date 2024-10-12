package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumRoot100Digit(t *testing.T) {
	result := SumRoot100Digit(2)
	assert.Equal(t, int64(475), result)
}
