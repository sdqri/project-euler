package main

import (
	"fmt"
	"math"
	. "math"
	"time"
)

func sieveOfEratosthenes(n int) []int {
	root := int(math.Sqrt(float64(n)))
	notPrimes := make([]bool, n)

	notPrimes[0] = true // 1 is not prime
	for i := 2; i <= root; i++ {
		factor := 2 * i
		for factor <= n {
			notPrimes[factor-1] = true
			factor += i
		}
	}

	estimatedNumberOfPrime := int64(float64(n) / Log(float64(n)))
	primes := make([]int, 0, estimatedNumberOfPrime)
	for i, isNotPrime := range notPrimes {
		if !isNotPrime {
			primes = append(primes, i+1)
		}
	}

	return primes
}

func main() {
	start := time.Now()
	n := 1_000_000
	primes := sieveOfEratosthenes(n)
	primesMap := make(map[int]struct{}, len(primes))
	for _, prime := range primes {
		primesMap[prime] = struct{}{}
	}

	MaxLength := 1
	maxSum := 2
	for i := range primes {
		currentLength := 0
		currentSum := 0
		for j := i; j < len(primes); j++ {
			currentSum += primes[j]
			currentLength += 1

			if _, ok := primesMap[currentSum]; ok && (currentLength > MaxLength) {
				MaxLength = currentLength
				maxSum = currentSum
			}
		}

	}
	elapsed := time.Since(start)
	format := "The longest sum of consecutive primes below %d that adds to a prime:\n" +
		"\tContains %d terms, and is equal to %d. (elapsed = %v)\n"
	fmt.Printf(format, n, MaxLength, maxSum, elapsed)
}
