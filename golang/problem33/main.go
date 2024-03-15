package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func IsDigitCancelling(numerator, denominator int, trivial bool) (bool, error) {
	strNumerator := fmt.Sprintf("%d", numerator)
	strDenominator := fmt.Sprintf("%d", denominator)
	if len(strNumerator) != 2 || len(strDenominator) != 2 {
		return false, nil
	}

	digits := make(map[rune]int)
	for _, r := range strNumerator {
		digits[r]++
	}
	for _, r := range strDenominator {
		digits[r]++
	}
	if len(digits) != 3 {
		return false, nil
	}

	var commonRune rune
	for digit, count := range digits {
		if count == 2 {
			commonRune = digit
		}
	}

	if !trivial {
		if commonRune == '0' {
			return false, nil
		}
	}

	simpleStrNumerator := strings.ReplaceAll(strNumerator, string(commonRune), "")
	simpleNumerator, err := strconv.Atoi(simpleStrNumerator)
	if err != nil {
		return false, err
	}
	simpleStrDenominator := strings.ReplaceAll(strDenominator, string(commonRune), "")
	simpleDenominator, err := strconv.Atoi(simpleStrDenominator)
	if err != nil {
		return false, err
	}

	if simpleDenominator == 0 {
		return false, nil
	}
	fraction := float64(numerator) / float64(denominator)
	simpleFraction := float64(simpleNumerator) / float64(simpleDenominator)
	if fraction != simpleFraction {
		return false, nil
	}

	return true, nil
}

func simplifyFraction(fraction Fraction) Fraction {
outerLoop:
	for {
		for i := 2; i < fraction.denominator; i++ {
			if fraction.numerator%i == 0 && fraction.denominator%i == 0 {
				fraction.numerator /= i
				fraction.denominator /= i
				continue outerLoop
			}
		}
		break
	}
	return fraction
}

type Fraction struct {
	numerator   int
	denominator int
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

func main() {
	start := time.Now()
	fractions := make([]Fraction, 0)
	for i := 1; i < 100; i++ {
		for j := i; j < 100; j++ {
			isDigitalCancelling, err := IsDigitCancelling(i, j, false)
			if isDigitalCancelling == true && err == nil {
				fractions = append(fractions, Fraction{numerator: i, denominator: j})
			}
		}
	}
	result := Fraction{1, 1}
	for _, fraction := range fractions {
		result.numerator *= fraction.numerator
		result.denominator *= fraction.denominator
	}
	elapsed := time.Since(start)

	fmt.Printf("fractions = %v\n", fractions)
	fmt.Printf("result = %v & simpleified version = %v (elapsed=%v)\n", result, simplifyFraction(result), elapsed)
}
