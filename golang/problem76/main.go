package main

import (
	"fmt"
	"time"
)

func P(k, n int64) int64 {
	if k > n {
		return 0
	}
	if k == 1 {
		return 1
	}
	return P(k-1, n-1) + P(k, n-k)
}

func main() {
	start := time.Now()
	ways := int64(0)
	for k := int64(2); k <= 100; k++ {
		ways += P(k, 100)
	}
	elapsed := time.Since(start)
	template := "Different ways which one hundred can be written as a sum of at least two positive integers = %v (elapsed = %v)\n"
	fmt.Printf(template, ways, elapsed)
}
