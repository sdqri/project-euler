package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func power(base, exponent int) *big.Int {
	bigBase := big.NewInt(int64(base))
	result := big.NewInt(int64(1))
	for range exponent {
		result.Mul(result, bigBase)
	}
	return result
}

func sumOfDigits(num *big.Int) int {
	result := 0
	digitsStr := num.String()
	for _, digitRune := range digitsStr {
		digit, err := strconv.Atoi(string(digitRune))
		if err != nil {
			panic(err)
		}
		result += digit
	}

	return result
}

func main() {
	start := time.Now()
	maxSumOfDigits := 1
	maxBase := 1
	maxExponent := 1
	for base := range 100 {
		for exponent := range 100 {
			if sumOfDigits(power(base, exponent)) > maxSumOfDigits {
				maxSumOfDigits = sumOfDigits(power(base, exponent))
				maxBase = base
				maxExponent = exponent
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("max sum of digits is for %d^%d = %d (elapsed : %v)\n",
		maxBase, maxExponent, maxSumOfDigits, elapsed)
}
