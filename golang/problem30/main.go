package main

import (
	"fmt"
	"time"
)

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func SumOfPowersOfDigits(number int, power int) int {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += IntPow(digit, power)
		number = number / 10
	}
	return sum
}

func main() {
	start := time.Now()
	numbers := make([]int, 0)
	for i := 2; i <= 1000000; i++ {
		if i == SumOfPowersOfDigits(i, 5) {
			numbers = append(numbers, i)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("all the numbers that can be written as the sum of fifth powers of their digits=%d (elapsed=%v)\n", sum(numbers), elapsed)
}
