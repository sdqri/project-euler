package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	rightTriangles := map[int64]map[string]struct{}{} // perimeter -> trianglesMap
	m := int64(1)
mLoop:
	for {
	nLoop:
		for n := int64(1); n < m; n++ {
			k := int64(1)
		kLoop:
			for {
				a := k * (m*m - n*n)
				b := k * (2 * m * n)
				c := k * (m*m + n*n)
				p := a + b + c

				if p > 1_500_000 {
					if k != 1 {
						break kLoop
					} else if n != 1 {
						break nLoop
					} else {
						break mLoop
					}
				}

				if a*a+b*b == c*c {
					sides := []int64{a, b, c}
					slices.Sort(sides)
					signature := fmt.Sprintf("%d, %d, %d", sides[0], sides[1], sides[2])
					_, ok := rightTriangles[p]
					if !ok {
						rightTriangles[p] = map[string]struct{}{}
					}
					rightTriangles[p][signature] = struct{}{}
				}

				k++
			}
		}
		m++
	}

	count := 0
	for _, v := range rightTriangles {
		if len(v) == 1 {
			count++
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("count L <= 1500000 that can form only one sided right angle triangle = %v (elapsed = %v)\n", count, elapsed)
}
