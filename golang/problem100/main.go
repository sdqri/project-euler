package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func FindBlueCounts(n int64) int64 {
	r := big.NewInt(1).Mul(big.NewInt(n), big.NewInt(n-1))
	b := int64(math.Ceil(float64(n) / math.Sqrt2))
	l := big.NewInt(1).Mul(big.NewInt(2*b), big.NewInt(b-1))
	if l.Cmp(r) == 0 {
		return b
	}
	return -1
}

func FindNCounts(b int64) int64 {
	n := int64(math.Floor(float64(b) * math.Sqrt2))
	r := big.NewInt(1).Mul(big.NewInt(n), big.NewInt(n-1))
	l := big.NewInt(1).Mul(big.NewInt(2*b), big.NewInt(b-1))
	if l.Cmp(r) == 0 {
		return n
	}
	return -1
}

func main() {
	start := time.Now()
	nSlice := []int64{4, 21}
	bSlice := []int64{3, 15}
outerLoop:
	for {
		l := len(nSlice)
		if nSlice[l-1] > 1e12 {
			break outerLoop
		}
		n := int64(float64(nSlice[l-1]) * (float64(nSlice[l-1]) / float64(nSlice[l-2])))
		bCounts := int64(-1)
		for {
			bCounts = FindBlueCounts(n)
			if bCounts != -1 {
				nSlice = append(nSlice, n)
				bSlice = append(bSlice, bCounts)
				break
			}
			n++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("first arrangement with n>=1e12 has %d number of blue discs (elapsed=%v)\n", bSlice[len(bSlice)-1], elapsed)
}
