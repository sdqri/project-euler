package main

import (
	"fmt"
	"strconv"
	"time"
)

func lenOfNumber(num int) int {
	return len(strconv.Itoa(num))
}

func champernownerConstantBase10(nth int) (int, error) {
	traversedLength := 0
	index := 1
	for traversedLength+lenOfNumber(index) < nth {
		traversedLength += lenOfNumber(index)
		index++
	}
	digit_idx := (nth - traversedLength) - 1
	digitRune := []rune(strconv.Itoa(index))[digit_idx]
	return strconv.Atoi(string(digitRune))
}

func main() {
	ns := []int{1, 10, 100, 1000, 10_000, 100_000, 1_000_000}

	start := time.Now()
	result := 1
	for _, n := range ns {
		digit, err := champernownerConstantBase10(n)
		if err != nil {
			panic(err)
		}
		result *= digit
	}
	elapsed := time.Since(start)

	fmt.Printf("d[1]*d[10]*d[100]*d[1000]*d[10_000]*d[100_000]*d[1_000_000] = %d (elapsed = %v)\n", result, elapsed)
}
