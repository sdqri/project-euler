package main

import (
	"fmt"
	"time"
)

func factorial(x int64) int64 {
	if x == 0 || x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func digitFactorial(x int64) int64 {
	digitFactorial := int64(0)
	for x > 0 {
		digit := x % 10
		digitFactorial += factorial(digit)
		x /= 10
	}
	return digitFactorial
}

func LenDigitFactorialChain(memory map[int64]int64, x int64) int64 {
	if memory != nil {
		l, ok := memory[x]
		if ok {
			return l
		}
	}
	numsMap := map[int64]struct{}{}
	numsMap[x] = struct{}{}
	chainLength := int64(1)
	next := x
outerLoop:
	for {
		next = digitFactorial(next)
		if _, ok := numsMap[next]; ok {
			break outerLoop
		}

		numsMap[next] = struct{}{}
		chainLength++
	}

	if memory != nil {
		memory[x] = chainLength
	}

	return chainLength
}

func main() {
	start := time.Now()
	mem := map[int64]int64{}
	count := 0
	for x := range int64(1e6) {
		if LenDigitFactorialChain(mem, x) == 60 {
			count++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Count x for x < 1e6 that (LenDigitFactorialChain(x) = 60) : %d (elapsed = %v)\n", count, elapsed)
}
