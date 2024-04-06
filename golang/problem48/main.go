package main

import (
	"fmt"
	"time"
)

func power(base int64, exponent int64) int64 {
	var result int64 = 1
	for i := 0; i < int(exponent); i++ {
		result *= base
	}
	return result
}

func powerWithModulo(base int64, exponent int64, modulo int64) int64 {
	var result int64 = 1
	for i := 0; i < int(exponent); i++ {
		result *= base
		result %= modulo
	}
	return result
}

func main() {
	start := time.Now()
	modulo := power(10, 10)
	sum := int64(0)
	for i := int64(1); i <= 1000; i++ {
		sum += powerWithModulo(i, i, modulo)
	}

	sum %= modulo
	elapsed := time.Since(start)
	fmt.Printf("sum = %d (elapsed=%v)\n", sum, elapsed)
}
