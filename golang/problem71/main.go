package main

import (
	"fmt"
	"math"
	"time"
)

var primeMap map[int64]bool

func IsPrime(x int64) bool {
	n := int64(math.Sqrt(float64(x)))
	for i := int64(2); i <= n; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

type Fraction struct {
	numerator   int64
	denominator int64
}

func (f Fraction) Value() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

func (f Fraction) Reduce() Fraction {
	gcd := GCD(f.numerator, f.denominator)
	return Fraction{
		numerator:   f.numerator / gcd,
		denominator: f.denominator / gcd,
	}
}

func (f Fraction) String() string {
	return fmt.Sprintf("%v/%v", f.numerator, f.denominator)
}

func New(numerator, denominator int64) Fraction {
	return Fraction{numerator, denominator}
}

func GCD(a, b int64) int64 {
	if b == 0 {
		return a
	}

	result := GCD(b, a%b)
	return result
}

func main() {
	start := time.Now()
	maxF := New(0, 1)
dLoop:
	for d := int64(1_000_000); d >= 1; d-- {
		for n := int64(1); n < d; n++ {
			numerator := 3*d - n
			denominator := 7 * d
			f := New(numerator, denominator)
			reducedF := f.Reduce()
			if reducedF.denominator <= 1_000_000 {
				if reducedF.Value() > maxF.Value() {
					maxF = reducedF
				}
				continue dLoop
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Println(maxF, elapsed)
}
