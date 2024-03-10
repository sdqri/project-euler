package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCombinations(t *testing.T) {
	aMin, aMax := 2, 5
	bMin, bMax := 2, 5

	expected := []string{"4", "8", "9", "16", "25", "27", "32", "64", "81", "125", "243", "256", "625", "1024", "3125"}
	sort.Sort(sort.StringSlice(expected))

	actual := GetDistinctCombinations(aMin, aMax, bMin, bMax)
	sort.Sort(sort.StringSlice(actual))
	require.Equal(
		t,
		actual,
		expected,
	)
}

func TestGetDinstictTerms(t *testing.T) {
	aMin, aMax := 2, 5
	bMin, bMax := 2, 5
	require.Equal(t, 15, len(GetDistinctCombinations(aMin, aMax, bMin, bMax)))
}
