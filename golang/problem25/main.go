package main

import (
	"fmt"
	"math/big"
	"time"
)

func Fibonacci(n int, memory map[int]*big.Int) *big.Int {
	if n == 1 || n == 2 {
		return big.NewInt(1)
	}
	if value, ok := memory[n]; ok {
		return value
	}
	result := big.NewInt(0).Add(Fibonacci(n-1, memory), Fibonacci(n-2, memory))
	memory[n] = result
	return result
}

func main() {
	memory := make(map[int]*big.Int)
	start := time.Now()
	i := 1
	for {
		result := fmt.Sprint(Fibonacci(i, memory))
		l := len(result)
		fmt.Println(l)
		if l == 1000 {
			break
		}
		i++
	}
	elapsed := time.Since(start)
	fmt.Printf("i=%d (elapsed time = %v)\n", i, elapsed)
}
