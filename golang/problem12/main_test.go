package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstTriangleNumberWithNDivisors(t *testing.T) {
	assert.Equal(t, int64(28), FindFirstTriangleNumberWithOverNDivisors(5))
}
