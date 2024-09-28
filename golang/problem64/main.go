package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	oddPeriodsCount := 0
	for i := 1; i <= 10000; i++ {
		if !isPerfectSquare(i) {
			if GetSquareRootPeriodLen(float64(i))%2 != 0 {
				oddPeriodsCount++
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Continued fractions for n<= 10000 that have an odd period = %d (elapsed = %v)\n", oddPeriodsCount, elapsed)
}

func isPerfectSquare(x int) bool {
	root := int(math.Floor(math.Sqrt(float64(x))))
	return root*root == x
}

func GetSquareRootPeriod(x float64) []int {
	loopMap := make(map[string]struct{})
	a := 0
	numerator := float64(0)
	denominator := float64(1)
	loopMap[fmt.Sprintf("%v-%v-%v", a, numerator, denominator)] = struct{}{}
	root := math.Sqrt(x)
	var series []int
	for {
		newA := math.Floor((root + float64(numerator)) / float64(denominator))
		newNuminator := newA*denominator - numerator
		newDenominator := (x - (newNuminator * newNuminator)) / denominator
		fmt.Println(fmt.Sprintf("%v-%v-%v", newA, newNuminator, newDenominator))
		if _, ok := loopMap[fmt.Sprintf("%v-%v-%v", newA, newNuminator, newDenominator)]; ok {
			break
		}
		series = append(series, int(newA))
		a = int(newA)
		numerator = newNuminator
		denominator = newDenominator
		loopMap[fmt.Sprintf("%v-%v-%v", a, numerator, denominator)] = struct{}{}
	}
	return series
}

func GetSquareRootPeriodLen(num float64) int {
	return len(GetSquareRootPeriod(num)) - 1
}
