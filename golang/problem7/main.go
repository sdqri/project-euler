package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(x int) bool {
	n := int(math.Sqrt(float64(x)))
	for i := 2; i <= n; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func GetNthPrime(N int) int {
	nth := 0
	x := 2
	for {
		if IsPrime(x) {
			if nth++; nth == N {
				break
			}
		}
		x++
	}
	return x
}

func main() {
	now := time.Now()
	fmt.Println(GetNthPrime(10001))
	elapsed := time.Since(now)
	fmt.Println(elapsed)
}
