package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	multiplier := int64(28433)
	power := 7830457
	exponentiationResult := int64(1)
	for range power {
		exponentiationResult *= 2
		exponentiationResult %= 1e11
	}
	result := multiplier*exponentiationResult + 1
	result %= 1e10
	elapsed := time.Since(start)
	fmt.Printf("last ten digits of %d * 2^%d + 1 = %d (elapsed = %v)\n", multiplier, power, result, elapsed)
}
