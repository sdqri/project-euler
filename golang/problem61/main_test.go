package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTriangleNumber(t *testing.T) {
	tests := []struct {
		input    int64
		expected bool
	}{
		{1, true},
		{3, true},
		{6, true},
		{10, true},
		{15, true},
		{7, false},
		{8, false},
	}

	for _, test := range tests {
		if result := isValidTriangleNumber(test.input); result != test.expected {
			t.Errorf("isValidTriangleNumber(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsValidSquareNumber(t *testing.T) {
	tests := []struct {
		input    int64
		expected bool
	}{
		{1, true},
		{4, true},
		{9, true},
		{16, true},
		{25, true},
		{2, false},
		{3, false},
		{-4, false},
	}

	for _, test := range tests {
		if result := isValidSquareNumber(test.input); result != test.expected {
			t.Errorf("isValidSquareNumber(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsValidPentagonalNumber(t *testing.T) {
	tests := []struct {
		input    int64
		expected bool
	}{
		{1, true},
		{5, true},
		{12, true},
		{22, true},
		{35, true},
		{7, false},
		{10, false},
	}

	for _, test := range tests {
		if result := isValidPentagonalNumber(test.input); result != test.expected {
			t.Errorf("isValidPentagonalNumber(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsValidHexagonalNumber(t *testing.T) {
	tests := []struct {
		input    int64
		expected bool
	}{
		{1, true},
		{6, true},
		{15, true},
		{28, true},
		{45, true},
		{7, false},
		{9, false},
	}

	for _, test := range tests {
		if result := isValidHexagonalNumber(test.input); result != test.expected {
			t.Errorf("isValidHexagonalNumber(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsValidHeptagonalNumber(t *testing.T) {
	tests := []struct {
		input    int64
		expected bool
	}{
		{1, true},
		{7, true},
		{18, true},
		{34, true},
		{55, true},
		{10, false},
		{20, false},
	}

	for _, test := range tests {
		if result := isValidHeptagonalNumber(test.input); result != test.expected {
			t.Errorf("isValidHeptagonalNumber(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsValidOctagonalNumber(t *testing.T) {
	tests := []struct {
		input    int64
		expected bool
	}{
		{1, true},
		{8, true},
		{21, true},
		{40, true},
		{65, true},
		{10, false},
		{18, false},
	}

	for _, test := range tests {
		if result := isValidOctagonalNumber(test.input); result != test.expected {
			t.Errorf("isValidOctagonalNumber(%d) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFindCyclesOfLength6(t *testing.T) {
	tests := []struct {
		graph    map[string][]string
		length   int
		expected [][]string
	}{
		{
			graph: map[string][]string{
				"A": {"B", "C", "D"},
				"B": {"A", "C", "E"},
				"C": {"A", "B", "F"},
				"D": {"A", "E", "F"},
				"E": {"B", "D", "F"},
				"F": {"C", "D", "E"},
			},
			length: 6,
			expected: [][]string{
				{"A", "B", "C", "F", "E", "D"},
				{"A", "B", "E", "D", "F", "C"},
				{"A", "C", "B", "E", "F", "D"},
				{"A", "C", "F", "D", "E", "B"},
				{"A", "D", "E", "F", "C", "B"},
				{"A", "D", "F", "E", "B", "C"},
			},
		},
	}
	for _, test := range tests {
		result := findCyclesWithLength(test.length, test.graph)
		assert.Len(t, result, len(test.expected), "number of cycles isn't equal to expected value")
	}
}
