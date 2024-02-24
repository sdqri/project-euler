package main

import (
	"fmt"
	"strings"
)

var onesAndTeens map[int]string = map[int]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens map[int]string = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var names map[int]string = map[int]string{
	2:  "hundred",
	3:  "thousand",
	6:  "million",
	9:  "billion",
	12: "trillion",
}

func numOfDigits(x int) int {
	num := 0
	for x > 0 {
		x /= 10
		num += 1
	}
	return num
}

func pow(x int, y int) int {
	result := 1
	for range y {
		result *= x
	}
	return result
}

func WriteNumbers(x int) string {
	num := numOfDigits(x) - 1
	if x >= 0 && x < 20 {
		return onesAndTeens[x]
	} else if x >= 20 && x < 100 {
		if x%10 == 0 {
			return tens[x/10]
		} else {
			return tens[x/10] + " " + WriteNumbers(x%10)
		}
	} else {
		var name string
		for {
			var ok bool
			name, ok = names[num]
			if ok {
				break
			}
			num--
		}
		result := make([]string, 0)
		result = append(result, WriteNumbers(x/pow(10, num)))
		result = append(result, name)
		if andNum := x % pow(10, num); andNum != 0 {
			result = append(result, "and")
			result = append(result, WriteNumbers(andNum))
		}
		return strings.Join(result, " ")
	}
}

func main() {
	l := 0
	for i := range 1000 {
		l += len(strings.ReplaceAll(WriteNumbers(i+1), " ", ""))
	}
	fmt.Println(l)

}
