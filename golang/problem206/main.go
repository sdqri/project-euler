package main

import (
	"fmt"
	"math"
	"time"
)

func MatchesPattern(num int64) bool {
	if (num/1e18)%10 != 1 {
		return false
	}

	if (num/1e16)%10 != 2 {
		return false
	}

	if (num/1e14)%10 != 3 {
		return false
	}

	if (num/1e12)%10 != 4 {
		return false
	}

	if (num/1e10)%10 != 5 {
		return false
	}

	if (num/1e8)%10 != 6 {
		return false
	}

	if (num/1e6)%10 != 7 {
		return false
	}

	if (num/1e4)%10 != 8 {
		return false
	}

	if (num/1e2)%10 != 9 {
		return false
	}

	if num%10 != 0 {
		return false
	}

	return true
}

func main() {
	start := time.Now()
	smalledsSquare := 1020304050607080900
	biggestSquare := 1929394959697989990

	smallestRoot := int64(math.Sqrt(float64(smalledsSquare)))
	biggestRoot := int64(math.Sqrt(float64(biggestSquare)))

	answer := int64(0)
	for root := smallestRoot; root <= biggestRoot; root++ {
		if MatchesPattern(root * root) {
			answer = root
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("integer whose square has the form 1_2_3_4_5_6_7_8_9_ = %d (elapsed = %v)\n", answer, elapsed)
}
