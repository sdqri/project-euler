package main

import (
	"fmt"
	"math"
	"time"
)

func CreateSpiral(size int) [][]int {
	values := make([][]int, 0, size)
	for i := 0; i < size; i++ {
		values = append(values, make([]int, size))
	}

	i, j := size-1, size-1
	iDelta, jDelta := 0, -1
	for n := size * size; n > 0; n-- {
		values[i][j] = n
		if i+iDelta < 0 || i+iDelta >= size ||
			j+jDelta < 0 || j+jDelta >= size || values[i+iDelta][j+jDelta] != 0 {
			if iDelta == 1 || iDelta == -1 {
				iDelta *= -1
			}
			iDelta, jDelta = jDelta, iDelta
		}
		i += iDelta
		j += jDelta
	}

	return values
}

func getSpiralPrimeRatio(spiral [][]int) float64 {
	size := len(spiral)
	primeCount := 0
	iMainDiag := 0
	iAntiDiag := size - 1
	for j := 0; j < size; j++ {
		if iMainDiag != iAntiDiag {
			if isPrime(spiral[iMainDiag][j]) {
				primeCount++
			}
			if isPrime(spiral[iAntiDiag][j]) {
				primeCount++
			}
		} else {
			if isPrime(spiral[iMainDiag][j]) {
				primeCount++
			}
		}
		iMainDiag += 1
		iAntiDiag -= 1
	}
	return float64(primeCount) / (2*float64(size) - 1)

}

func isPrime(x int) bool {
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

func main() {
	start := time.Now()
	n := 5
	primeCount := 3
	for {
		e4 := n * n
		e3 := e4 - (n - 1)
		e2 := e3 - (n - 1)
		e1 := e2 - (n - 1)
		if isPrime(e4) {
			primeCount++
		}
		if isPrime(e3) {
			primeCount++
		}
		if isPrime(e2) {
			primeCount++
		}
		if isPrime(e1) {
			primeCount++
		}
		spiralPrimeRatio := float64(primeCount) / (2*float64(n) - 1)
		fmt.Printf("n=%d, prime ratio=%f\n", n, spiralPrimeRatio)
		if spiralPrimeRatio < 0.1 {
			break
		}
		n += 2
	}

	elapsed := time.Since(start)
	fmt.Printf("side length of the square spiral for which the ratio of primes along both diagonals first falls below = %d (elapsed : %v)\n",
		n, elapsed)

}
