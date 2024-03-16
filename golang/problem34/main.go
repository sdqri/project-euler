package main

import (
	"fmt"
	"time"
)

func factorial(n int) int {
	factorialMap := map[int]int{
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
		6: 720,
		7: 5040,
		8: 40320,
		9: 362880,
	}

	if value, ok := factorialMap[n]; ok {
		return value
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func DigitFactorial(x int) int {
	result := 0
	for {
		// fmt.Println("x=", x)
		// fmt.Println("digit", x%10)
		// fmt.Println("divisor", 10)
		if x/10 <= 0 {
			if x > 0 {
				result += factorial(x % 10)
			}
			break
		}
		result += factorial(x % 10)
		x /= 10
	}
	return result
}

func main() {
	start := time.Now()
	sum := 0
	for i := 3; i <= 10_000_000; i++ {
		if DigitFactorial(i) == i {
			sum += i
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("sum=%v (elapsed=%v)\n", sum, elapsed)
}
