package main

import (
	"fmt"
	"math/big"
	"strings"
)

func getUnitFraction(n int) *big.Float {
	bigOne := big.NewFloat(1).SetPrec(20000)
	bigN := big.NewFloat(float64(n)).SetPrec(20000)
	bigOne.Quo(bigOne, bigN)
	return bigOne
}

func GetLargestRecurringCycle(s string) string {
	for i := range len(s) / 2 {
		recurringCycle := make([]string, 0)
		index := i
		for {
			remainedSlice := make([]string, 0)
			parts := strings.Split(
				s,
				strings.Join(recurringCycle, ""),
			)
			remainedSlice = append(remainedSlice, parts[0])
			for i := 1; i <= len(parts)-1; i++ {
				if parts[i] == "" {
					continue
				} else if !strings.HasSuffix(strings.Join(recurringCycle, ""), parts[i]) {
					remainedSlice = append(remainedSlice, parts[i])
				}
			}
			remainedString := strings.Join(remainedSlice, "")

			if len(remainedString) == i &&
				strings.Count(s, strings.Join(recurringCycle, "")) > 1 {
				return strings.Join(recurringCycle, "")
			}
			recurringCycle = append(recurringCycle, string(s[index]))
			index++
			if index == len(s)-1 {
				break
			}
		}
	}
	return ""
}

func main() {
	d := -1
	max := 0
	bigCycle := ""
	for i := range 1000 - 1 {
		fmt.Println(i)
		unitFraction := getUnitFraction(i + 1)
		recurringCycle := GetLargestRecurringCycle(unitFraction.Text('f', 2000))
		if len(recurringCycle) > max {
			max = len(recurringCycle)
			bigCycle = recurringCycle
			d = i + 1
		}
	}

	fmt.Println(max, bigCycle, d)
}
