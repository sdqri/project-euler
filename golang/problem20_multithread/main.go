package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

func factorial(start int, end int) *big.Int {
	result := big.NewInt(1)
	for i := start; i >= end; i-- {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func parallelFactorial(x int, parallelism int) *big.Int {
	resultChan := make(chan *big.Int)
	size := x / parallelism
	for i := 0; i < parallelism; i++ {
		var start int
		end := size*i + 1
		if i == parallelism-1 {
			start = x
		} else {
			start = size*i + size
		}
		go computeFactorial(start, end, resultChan)
	}

	middleResults := make([]*big.Int, 0)
	for i := 0; i < parallelism; i++ {
		middleResults = append(middleResults, <-resultChan)
	}
	return parallelMultiply(middleResults)
}

func computeFactorial(start int, end int, resultChan chan *big.Int) {
	resultChan <- factorial(start, end)
}

func SumOfDigits(x *big.Int) int {
	sum := 0
	for _, char := range x.String() {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}
	return sum
}

func parallelMultiply(inputs []*big.Int) *big.Int {
	rand.Shuffle(len(inputs), func(i, j int) {
		inputs[i], inputs[j] = inputs[j], inputs[i]
	})
	if len(inputs) == 1 {
		return inputs[0]
	}
	resultChan := make(chan *big.Int)
	for {
		results := make([]*big.Int, 0)
		n := len(inputs) / 2
		for i := range n {
			go mul(inputs[i*2], inputs[i*2+1], resultChan)
		}
		for range n {
			results = append(results, <-resultChan)
		}
		if len(inputs)%2 != 0 {
			results = append(results, inputs[len(inputs)-1])
		}
		if len(results) == 1 {
			return results[0]
		}
		inputs = results
	}
}

func mul(x *big.Int, y *big.Int, resultChan chan *big.Int) {
	x = x.Mul(x, y)
	resultChan <- x
	return
}

func main() {
	x := 100
	start := time.Now()
	result := SumOfDigits(parallelFactorial(100, 1))
	elapsed := time.Since(start)
	fmt.Printf("sum of digits for %v! = %v (elapsed time = %v)\n", x, result, elapsed)
}
