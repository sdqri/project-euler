package main

import (
	"fmt"
	"math"
	"time"
)

func is_prime(x int) bool {
	if x < -1 {
		x *= -1
	}
	root := math.Sqrt(float64(x))
	if x == 0 || x == 1 {
		return false
	}
	for i := 2; i <= int(root); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

type QuadraticFormula struct {
	a int
	b int
}

func (qf QuadraticFormula) GetValue(n int) int {
	return int(math.Pow(float64(n), 2)) + qf.a*n + qf.b
}

func getScore(quadraticFormula QuadraticFormula) int {
	n := 0
	for {
		value := quadraticFormula.GetValue(n)
		if !is_prime(value) {
			return n - 1
		}
		n++
	}
}

func main() {
	start := time.Now()
	maxScore := 0
	maxA := 0
	maxB := 0

	for b := -1000; b <= 1000; b++ {
		if !is_prime(b) {
			continue
		}
		for a := -999; a < 1000; a++ {
			quadraticFormula := QuadraticFormula{a: a, b: b}
			score := getScore(quadraticFormula)
			if score >= maxScore {
				maxScore = score
				maxA = a
				maxB = b
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("max prime = %d for n^2+%dn+%d (elapsed = %v)\n", maxScore, maxA, maxB, elapsed)
	fmt.Printf("a*b = %d\n", maxA*maxB)
}
