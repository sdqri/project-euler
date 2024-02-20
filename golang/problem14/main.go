package main

import (
	"fmt"
	"time"
)

func calc_collatz_chain_length(memory map[int]int, x int) int {
	if x == 1 {
		return 1
	} else if l, ok := memory[x]; ok {
		return l
	} else {
		l := 0
		if x%2 == 0 {
			l = calc_collatz_chain_length(memory, x/2) + 1
		} else {
			l = calc_collatz_chain_length(memory, 3*x+1) + 1
		}
		memory[x] = l
		return l
	}
}

func main() {
	start := time.Now()
	memory := make(map[int]int)
	var max_length, max_x int
	for x := range 1000000 {
		x++
		length := calc_collatz_chain_length(memory, x)
		if length > max_length {
			max_x = x
			max_length = length
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("max_length=%v & max_x=%v (time=%v)\n", max_length, max_x, elapsed/time.Millisecond)
}
