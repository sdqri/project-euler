package main

import (
	"fmt"
	"time"
)

func getPerimeter(n int) int {
	if n == 1 {
		return 1
	}
	return n*n - (n-2)*(n-2)
}

func SumOfDiagonalsInSpiral(size int) int {
	values := make([]int, 0)
	for i := range size * size {
		values = append(values, i)
	}

	sumOfDiagonals := 0
	for n := 1; n <= size; n += 2 {
		if n == 1 {
			sumOfDiagonals += 1
			continue
		}
		perimeter := getPerimeter(n)
		stride := (perimeter - 4) / 4
		i1 := stride + 1 + (n-2)*(n-2)
		i2 := i1 + stride + 1
		i3 := i2 + stride + 1
		i4 := i3 + stride + 1

		sumOfDiagonals = sumOfDiagonals + i1 + i2 + i3 + i4
	}

	return sumOfDiagonals
}

func main() {
	start := time.Now()
	n := 1001
	sumOfDiagonals := SumOfDiagonalsInSpiral(n)
	elapsed := time.Since(start)
	fmt.Printf("Sum of diagonals in a %d*%d size = %d (elapsed = %v)\n", n, n, sumOfDiagonals, elapsed)
}
