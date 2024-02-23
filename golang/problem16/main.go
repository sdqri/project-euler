package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	powerResult := big.NewInt(1)
	for range 1000 {
		powerResult.Mul(powerResult, big.NewInt(2))
	}
	powerResultString := powerResult.String()
	result := 0
	for _, rune := range powerResultString {
		digit, err := strconv.Atoi(string(rune))
		if err != nil {
			panic(err)
		}
		result += digit
	}
	elapsed := time.Since(start)
	fmt.Printf(
		"sum of digits for 2**%v = %v  (elapsed time = %v)\n",
		1000,
		result,
		elapsed,
	)

}
