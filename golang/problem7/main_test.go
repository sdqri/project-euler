package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNthPrime(t *testing.T) {
	n6 := GetNthPrime(6)
	assert.Equal(t, n6, 13, "6th prime should be 13")
}
