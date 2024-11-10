package main

import (
	"fmt"
	"math"
	"time"
)

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

type Fraction struct {
	Numerator   int64
	Denominator int64
}

func New(n, d int64) Fraction {
	return Fraction{n, d}
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

func (f Fraction) Reduce() Fraction {
	v := gcd(f.Numerator, f.Denominator)
	if v == 1 {
		return f
	}
	return New(f.Numerator/v, f.Denominator/v)
}

func main() {
	start := time.Now()
	fractions := map[Fraction]struct{}{}
	for d := int64(12000); d > 1; d-- {
		startN := int64(float64(d) / 3)
		if startN == d/3 {
			startN += 1
		}
		endN := int64(math.Ceil(float64(d) / 2))
		for n := int64(startN); n < endN; n++ {
			reducedFraction := New(n, d).Reduce()
			fractions[reducedFraction] = struct{}{}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("fractions between 1/3 and 1/2 (d<= 12000) = %v (elapsed = %v)", len(fractions), elapsed)
}
