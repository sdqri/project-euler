package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

func factorize(x int64) []int64 {
	if x <= 1 {
		return []int64{}
	}

	var factors []int64
	for x%2 == 0 {
		factors = append(factors, 2)
		x /= 2
	}

	for i := int64(3); i*i <= x; i += 2 {
		for x%i == 0 {
			factors = append(factors, i)
			x /= i
		}
	}

	if x > 2 {
		factors = append(factors, x)
	}

	return factors
}

func getUniqueFactors(x int64) []int64 {
	factors := factorize(x)
	uniqueFactorsMap := map[int64]struct{}{}
	for _, f := range factors {
		uniqueFactorsMap[f] = struct{}{}
	}

	uniqueFactors := []int64{}
	for f := range uniqueFactorsMap {
		uniqueFactors = append(uniqueFactors, f)
	}

	return uniqueFactors
}

func phi(x int64) int64 {
	result := float64(x)
	uniqueFactors := getUniqueFactors(x)
	for _, uf := range uniqueFactors {
		result *= (1 - (1 / float64(uf)))
	}

	return int64(result)
}

func isPermutation(x, y int64) bool {
	xDigitsSlice := strings.Split(fmt.Sprintf("%d", x), "")
	sort.Sort(sort.StringSlice(xDigitsSlice))
	xStr := strings.Join(xDigitsSlice, "")
	yDigitsSlice := strings.Split(fmt.Sprintf("%d", y), "")
	sort.Sort(sort.StringSlice(yDigitsSlice))
	yStr := strings.Join(yDigitsSlice, "")
	return xStr == yStr
}

func main() {
	start := time.Now()
	minValue := float64(math.MaxFloat64)
	n := int64(-1)
	for i := int64(2); i < 1e7; i++ {
		p := phi(i)
		val := float64(i) / float64(p)
		if isPermutation(i, phi(i)) && val < minValue {
			minValue = val
			n = i
		}
	}
	template := "Value of n for minimum n/phi(n) in which phi(n) is a permutation of n = %v (elapsed = %v)\n"
	elapsed := time.Since(start)
	fmt.Printf(template, n, elapsed)
}
