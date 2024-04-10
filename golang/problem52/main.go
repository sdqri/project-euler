package main

import (
	"fmt"
	"strings"
	"time"
)

func getDigits(n int) [10]bool {
	var digits [10]bool
	nStr := fmt.Sprintf("%d", n)
	for i := 0; i < 10; i++ {
		if strings.Count(nStr, fmt.Sprintf("%d", i)) > 0 {
			digits[i] = true
		}
	}

	return digits
}

func main() {
	start := time.Now()
	i := 1
	for {
		digits := getDigits(i)
		if digits == getDigits(i*2) &&
			digits == getDigits(i*3) &&
			digits == getDigits(i*4) &&
			digits == getDigits(i*5) &&
			digits == getDigits(i*6) {
			break
		}
		i++
	}
	elapsed := time.Since(start)
	fmt.Printf("Smallest x, that has Permuted Multiple of 2x, 3x, 4x, 5x, 6x =%d (elapsed=%v)\n", i, elapsed)
}
