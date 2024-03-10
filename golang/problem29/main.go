package main

import (
	"fmt"
	"math/big"
	"time"
)

func power(a int, b int) big.Int {
	result := big.NewInt(1)
	for range a {
		result.Mul(result, big.NewInt(int64(b)))
	}
	return *result
}

func GetDistinctCombinations(aMin, aMax, bMin, bMax int) []string {
	combinations := make([]big.Int, 0)
	for a := aMin; a <= aMax; a++ {
		for b := bMin; b <= bMax; b++ {
			combinations = append(combinations, power(a, b))
		}
	}

	uniqueTerms := make(map[string]struct{})
	for _, combination := range combinations {
		uniqueTerms[combination.Text(10)] = struct{}{}
	}

	uniqueCombinations := make([]string, 0)
	for uniqueTerm := range uniqueTerms {
		uniqueCombinations = append(uniqueCombinations, uniqueTerm)
	}

	return uniqueCombinations
}

func main() {
	start := time.Now()
	aMin, aMax := 2, 100
	bMin, bMax := 2, 100
	distinctTerms := len(GetDistinctCombinations(aMin, aMax, bMin, bMax))
	elapsed := time.Since(start)
	fmt.Printf(
		"distict terms for %d<=a<=%d and %d<=b<=%d = %d (elapsed = %v)\n",
		aMin, aMax,
		bMin, bMax,
		distinctTerms, elapsed,
	)
}
