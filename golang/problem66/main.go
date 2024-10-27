package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

func isPerfectSquare(x int) bool {
	root := int(math.Floor(math.Sqrt(float64(x))))
	return root*root == x
}

func GetSquareRootPeriod(x float64) []int64 {
	loopMap := make(map[string]struct{})
	a := 0
	numerator := float64(0)
	denominator := float64(1)
	loopMap[fmt.Sprintf("%v-%v-%v", a, numerator, denominator)] = struct{}{}
	root := math.Sqrt(x)
	var series []int64
	for {
		newA := math.Floor((root + float64(numerator)) / float64(denominator))
		newNuminator := newA*denominator - numerator
		newDenominator := (x - (newNuminator * newNuminator)) / denominator
		if _, ok := loopMap[fmt.Sprintf("%v-%v-%v", newA, newNuminator, newDenominator)]; ok {
			break
		}
		series = append(series, int64(newA))
		a = int(newA)
		numerator = newNuminator
		denominator = newDenominator
		loopMap[fmt.Sprintf("%v-%v-%v", a, numerator, denominator)] = struct{}{}
	}
	return series
}

var mem map[int64][]int64

func GetNthElement(x int64, n int64) int64 {
	period, ok := mem[x]
	if !ok {
		period = GetSquareRootPeriod(float64(x))
		mem[x] = period
	}

	if n < int64(len(period)-1) {
		return period[n]

	}

	i := ((n - 1) % int64(len(period)-1)) + 1

	return period[i]
}

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

func CalcNthConvergent(x int64, n int64) Fraction {
	convergent := New(GetNthElement(x, 0), 1)
	N := n

	var f func(int64) Fraction
	f = func(n int64) Fraction {
		if n == 1 {
			return New(0, 1)
		}

		i := N - n
		return Fraction{Numerator: big.NewInt(GetNthElement(x, i+1)), Denominator: big.NewInt(1)}.
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

func SolvePellsEquation(D int64) (*big.Int, *big.Int) {
	n := int64(1) // start with the first convergent
	for {
		f := CalcNthConvergent(D, n)
		// Calculate x^2 - D * y^2
		// if f.Numerator.Cmp(big.NewInt(math.MaxInt)) == 1 {
		// 	return -1, -1
		// }

		check := new(big.Int).Sub(
			new(big.Int).Mul(f.Numerator, f.Numerator), // x^2
			new(big.Int).Mul(big.NewInt(D), // D * y^2
				new(big.Int).Mul(f.Denominator, f.Denominator)),
		)

		if check.Cmp(big.NewInt(1)) == 0 { // if x^2 - D*y^2 == 1
			return f.Numerator, f.Denominator
		}

		n++
	}
}

func main() {
	start := time.Now()
	mem = map[int64][]int64{}
	biggestX := big.NewInt(-1)
	biggestD := int64(-1)
	for d := int64(1); d <= 1000; d++ {
		if int64(math.Sqrt(float64(d)))*int64(math.Sqrt(float64(d))) == d {
			continue
		}
		x, _ := SolvePellsEquation(d)
		if x.Cmp(biggestX) == 1 {
			biggestX = x
			biggestD = d
		}
	}
	elapsed := time.Since(start)
	template := "Find the value of D<=1000 in minimial solution of x for which the largest value of x is obtained = %v (elapsed = %v)\n"
	fmt.Printf(template, biggestD, elapsed)
}
