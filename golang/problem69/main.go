package main

import (
	"fmt"
	"time"
)

func nOverPhi(n int64) float64 {
	return float64(n) / float64(phi(n))
}

func factorize(n int64) []int64 {
	var factors []int64

	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for i := int64(3); i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func uniqueFactorize(n int64) []int64 {
	factorsSet := map[int64]struct{}{}
	for _, f := range factorize(n) {
		factorsSet[f] = struct{}{}
	}

	factors := []int64{}
	for f := range factorsSet {
		factors = append(factors, f)
	}
	return factors
}

func phi(n int64) int64 {
	factors := uniqueFactorize(n)
	p := float64(n)
	for _, f := range factors {
		p *= 1 - 1/float64(f)
	}
	return int64(p)
}

func main() {
	start := time.Now()
	maxNOverPhi := float64(0)
	maxN := int64(0)
	for i := int64(2); i <= 1_000_000; i++ {
		fmt.Println(i, " = ", phi(i))
		nOverP := nOverPhi(i)
		if nOverP > float64(maxNOverPhi) {
			maxNOverPhi = nOverP
			maxN = i
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("maximum n/phi(n) is %v/phi(%v) = %v (elapsed = %v)\n", maxN, maxN, maxNOverPhi, elapsed)
}
