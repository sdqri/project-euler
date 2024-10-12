package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func isPerfectSquare(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return root*root == num
}

func SumRoot100Digit(num int) int64 {
	bigNum := big.NewFloat(float64(num)).SetPrec(100 * 4)
	root := new(big.Float).Sqrt(bigNum)

	resultStr := fmt.Sprintf("%.110f", root)

	resultStr = strings.Replace(resultStr, ".", "", 1)

	first100Digits := resultStr[:100]

	sum := int64(0)
	for _, digitChar := range first100Digits {
		digit, err := strconv.ParseInt(string(digitChar), 10, 64)
		if err != nil {
			panic(err)
		}
		sum += digit
	}

	return sum
}

func main() {
	start := time.Now()
	totalSum := int64(0)

	for naturalNum := 1; naturalNum <= 100; naturalNum++ {
		if !isPerfectSquare(naturalNum) {
			totalSum += SumRoot100Digit(naturalNum)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("The total of the digital sums of the first one hundred decimal digits for all the irrational square roots\n")
	fmt.Printf("For the first one hundred natural numbers = %v (elapsed = %v)\n", totalSum, elapsed)
}
