package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func factorial(x int) *big.Int {
	if x == 1 {
		return big.NewInt(1)
	}
	return big.NewInt(1).Mul(big.NewInt(int64(x)), factorial(x-1))
}

func SumOfDigits(x *big.Int) int {
	sum := 0
	for _, char := range x.String() {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	return sum
}

func main() {
	x := 100
	start := time.Now()
	result := SumOfDigits(factorial(x))
	elapsed := time.Since(start)
	fmt.Printf("sum of digits for %v! = %v (elapsed time = %v)\n", x, result, elapsed)
}
