package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func gcd(a, b *big.Int) *big.Int {
	if a.Sign() == 0 {
		return new(big.Int).Set(b)
	}
	return gcd(new(big.Int).Mod(b, a), a)
}

func lcm(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(new(big.Int).Div(a, gcd(a, b)), b)
}

type Fraction struct {
	Numerator   *big.Int
	Denominator *big.Int
}

func New(num, den int64) Fraction {
	return Fraction{
		Numerator:   big.NewInt(num),
		Denominator: big.NewInt(den),
	}
}

func (f Fraction) String() string {
	return fmt.Sprintf("%v/%v", f.Numerator, f.Denominator)
}

func (f Fraction) Add(operand Fraction) Fraction {
	denominator := lcm(f.Denominator, operand.Denominator)
	numerator := new(big.Int).Add(
		new(big.Int).Mul(new(big.Int).Div(denominator, f.Denominator), f.Numerator),
		new(big.Int).Mul(new(big.Int).Div(denominator, operand.Denominator), operand.Numerator),
	)
	return Fraction{Numerator: numerator, Denominator: denominator}
}

func (f Fraction) Reciprocal() Fraction {
	return Fraction{Numerator: f.Denominator, Denominator: f.Numerator}
}

func GetNthElement(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(2)
	}

	if n%3 == 2 {
		return big.NewInt(2 * ((n / 3) + 1))
	}

	return big.NewInt(1)
}

func CalcNthConvergent(n int64) Fraction {
	convergent := New(2, 1)
	N := n

	var f func(int64) Fraction
	f = func(n int64) Fraction {
		if n == 1 {
			return New(0, 1)
		}

		i := N - n
		return Fraction{Numerator: GetNthElement(i + 1), Denominator: big.NewInt(1)}.
			Add(f(n - 1)).
			Reciprocal()
	}

	return convergent.Add(f(n))
}

func SumOfDigits(num big.Int) int64 {
	sum := int64(0)
	strNum := num.String()
	for _, c := range []rune(strNum) {
		digit, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			panic(err)
		}
		sum += digit
	}

	return sum
}

func main() {
	start := time.Now()
	result := CalcNthConvergent(100)
	sumOfNumerator := SumOfDigits(*result.Numerator)
	elapsed := time.Since(start)
	template := "100th convergent of the continued fraction for e = %v (elapsed = %v)\n"
	fmt.Printf(template, sumOfNumerator, elapsed)
}
