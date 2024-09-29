package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetExponentInBase10(base, exponent int64) float64 {
	multiplier := math.Log(float64(base))
	return multiplier * float64(exponent)
}

func main() {
	f, err := os.Open("./0099_base_exp.txt")
	if err != nil {
		panic(err)
	}
	start := time.Now()
	scanner := bufio.NewScanner(f)
	maxExponentInBase10 := float64(0)
	maxBase := int64(0)
	maxExponent := int64(0)
	maxLineNumber := 0
	lineNumber := 1
	for scanner.Scan() {
		lineText := scanner.Text()
		columns := strings.Split(lineText, ",")
		baseStr := strings.Trim(columns[0], " ")
		base, err := strconv.ParseInt(baseStr, 10, 64)
		if err != nil {
			panic(err)
		}
		exponentStr := strings.Trim(columns[1], " ")
		exponent, err := strconv.ParseInt(exponentStr, 10, 64)
		if err != nil {
			panic(err)
		}
		exponentInBase10 := GetExponentInBase10(base, exponent)
		if exponentInBase10 > maxExponentInBase10 {
			maxExponentInBase10 = exponentInBase10
			maxBase = base
			maxExponent = exponent
			maxLineNumber = lineNumber
		}
		lineNumber++
	}

	elapsed := time.Since(start)
	fmt.Printf("lineNumber=%v, base=%v, exponent=%v (elapsed=%v)\n", maxLineNumber, maxBase, maxExponent, elapsed)
}
