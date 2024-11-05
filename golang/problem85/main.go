package main

import (
	"fmt"
	"math"
	"time"
)

func getTriangleNumber(n int64) int64 {
	return (n * (n + 1)) / 2
}

func countRectangles(w, h int64) int64 {
	return getTriangleNumber(w) * getTriangleNumber(h)
}

func main() {
	start := time.Now()
	nearestNeightbor := int64(math.MaxInt64)
	nearestNeighborWidth, nearestNeighborHeight := int64(0), int64(0)
	w := int64(1)
widthLoop:
	for {
		h := int64(1)
	hightLoop:
		for {
			c := countRectangles(w, h)
			delta := math.Abs(float64(c - 2_000_000))
			if delta < math.Abs(float64(nearestNeightbor-2_000_000)) {
				nearestNeightbor = c
				nearestNeighborWidth, nearestNeighborHeight = w, h
			}

			if c > 2_000_000 {
				if h == 1 {
					break widthLoop
				}
				break hightLoop
			}

			h++
		}
		w++
	}
	elapsed := time.Since(start)
	template := "The area of the grid with the nearest solution = %v (elapsed = %v)\n"
	fmt.Printf(template, nearestNeighborWidth*nearestNeighborHeight, elapsed)
}
