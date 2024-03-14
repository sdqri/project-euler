package main

import (
	"fmt"
	"strings"
	"time"
)

func isPandigital(a, b int) bool {
	digit := make(map[rune]struct{})
	aStr := fmt.Sprintf("%d", a)
	bStr := fmt.Sprintf("%d", b)
	cStr := fmt.Sprintf("%d", a*b)
	l := len(aStr) + len(bStr) + len(cStr)
	if l != 9 {
		return false
	}
	aStr = strings.ReplaceAll(aStr, "0", "")
	bStr = strings.ReplaceAll(bStr, "0", "")
	cStr = strings.ReplaceAll(cStr, "0", "")

	for _, i := range aStr {
		digit[i] = struct{}{}
	}

	for _, i := range bStr {
		digit[i] = struct{}{}
	}

	for _, i := range cStr {
		digit[i] = struct{}{}
	}

	if len(digit) == 9 {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	products := make(map[int]struct{})

	for i := range 10000 {
		for j := range 100 {
			_, ok := products[i*j]
			if isPandigital(i, j) && !ok {
				products[i*j] = struct{}{}
			}
		}
	}

	sum := 0
	for key := range products {
		sum += key
	}
	elapsed := time.Since(start)
	fmt.Printf("sum of pandigital products=%d (elapsed=%v)\n", sum, elapsed)

}
