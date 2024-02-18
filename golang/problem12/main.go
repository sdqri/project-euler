package main

import (
	"fmt"
	"math"
	"time"
)

func CountDivisors(x int64) int64 {
	var n int64 = 0
	root := math.Sqrt(float64(x))
	rootFloor := int64(math.Floor(root))
	for divisor := range rootFloor {
		if x%(divisor+1) == 0 {
			n += 2
		}
	}
	if root == float64(rootFloor) {
		n--
	}
	return n
}

func FindFirstTriangleNumberWithOverNDivisors(nDivisor int64) int64 {
	var number int64 = 0
	var triangleNumber int64 = 0
	for {
		number = number + 1
		triangleNumber = triangleNumber + number
		// fmt.Printf("number=%d, triangleNumber=%d\n", number, triangleNumber)
		if CountDivisors(triangleNumber) >= nDivisor {
			break
		}
	}
	return triangleNumber
}

func main() {
	start := time.Now()
	triangleNumber := FindFirstTriangleNumberWithOverNDivisors(500)
	elapsed := time.Since(start)
	fmt.Printf("triangle number = %v, elapsed time = %v ms", triangleNumber, int64(elapsed/time.Millisecond))
}
