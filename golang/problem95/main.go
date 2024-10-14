package main

import (
	"fmt"
	"math"
	"time"
)

func GetProperDivisors(num int64) []int64 {
	if num <= 1 {
		return nil
	}

	divisors := []int64{1}
	sqrt := int64(math.Sqrt(float64(num)))

	for i := int64(2); i <= sqrt; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			if i != num/i { // To avoid adding the square root twice for perfect squares
				divisors = append(divisors, num/i)
			}
		}
	}

	return divisors
}

func SumOfProperDivisors(num int64) int64 {
	divisors := GetProperDivisors(num)
	sum := int64(0)
	for _, divisor := range divisors {
		sum += divisor
	}

	return sum
}

func LenAmicableChains(mem map[int64]int64, num int64, limit int64) int64 {
	l, ok := mem[num]
	if ok {
		return l
	}

	chainMap := map[int64]struct{}{}
	chain := []int64{num}

	next := num
loop:
	for {
		l, ok := mem[num]
		if ok {
			for i, n := range chain {
				mem[n] = int64(len(chain)-(i+1)) + l
			}
			return int64(len(chain)) + l
		}
		next = SumOfProperDivisors(next)
		if next > limit {
			return -1
		}

		if next == num {
			break loop
		}

		if _, ok := chainMap[next]; ok {
			return -1
		}

		chain = append(chain, next)
		chainMap[next] = struct{}{}
	}

	for i, n := range chain {
		mem[n] = int64(len(chain) - (i + 1))
	}
	return int64(len(chain))
}

func main() {
	start := time.Now()
	maxLen := int64(0)
	maxStart := int64(0)
	mem := map[int64]int64{}
	for x := range 1_000_000 {
		l := LenAmicableChains(mem, int64(x), 1_000_000)
		if l > int64(maxLen) {
			maxLen = l
			maxStart = int64(x)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("longest amicable chain with no element exceeding one million = %v with chain of %v numbers (elapsed = %v)\n", maxStart, maxLen, elapsed)
}
