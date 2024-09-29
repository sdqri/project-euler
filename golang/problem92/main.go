package main

import (
	"fmt"
	"time"
)

func squareDigit(x int64) int64 {
	result := int64(0)
	for x != 0 {
		remainder := x % 10
		result += remainder * remainder
		x /= 10
	}
	return result
}

var memory map[int64]int64

func endSquareDigitChain(x int64) int64 {
	if end, ok := memory[x]; ok {
		return end
	}

	tempSlice := []int64{x}
	next := x
	for {
		next = squareDigit(next)
		if next == 1 || next == 89 {
			break
		}
		tempSlice = append(tempSlice, next)
	}

	for _, e := range tempSlice {
		memory[e] = next
	}
	return next
}

func main() {
	memory = make(map[int64]int64)
	start := time.Now()
	endsIn89 := 0
	for i := int64(1); i <= 1e7; i++ {
		end := endSquareDigitChain(i)
		if end == 89 {
			endsIn89++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("numbers below ten million that their square digit chain will arrive at 89 = %v (elapsed = %v)\n", endsIn89, elapsed)
}
