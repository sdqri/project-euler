package main

import (
	"fmt"
	"strconv"
	"strings"
)

func arrayToStr(a []int) string {
	s := make([]string, 0)
	for _, e := range a {
		s = append(s, strconv.Itoa(e))
	}
	return strings.Join(s, "")
}

func endsIn(src []int, pattern []int) (bool, int) {
	fmt.Println("@@@@@@@@@@@@")
	srcStr := arrayToStr(src)
	patternStr := arrayToStr(pattern)
	for i := 0; i < len(src)-len(patternStr); i++ {
		remainderStr := strings.ReplaceAll(srcStr, patternStr, "")
		if strings.HasPrefix(patternStr, remainderStr) {
			fmt.Println("i=", i, "patternStr", patternStr, "reaminder=", remainderStr)
			return true, i
		}
		srcStr = srcStr[1:]
	}
	fmt.Println("@@@@@@@@@@@@")
	return false, len(src)-1
}

func CutRepeatingPattern(nums []int) ([]int, []int, error) {
	minStartIndex := len(nums) - 1
	minPattern := []int{}

	for i := 0; i <= len(nums)-1; i++ {
		pattern := []int{}
		for j := i; j <= len(nums)-1; j++ {
			pattern = append(pattern, nums[j])
			fmt.Println("-----------------------------")
			fmt.Println("pattern=", pattern)
			ends, startIndex := endsIn(nums, pattern)
			fmt.Println("|", ends, startIndex)
			if ends && startIndex < minStartIndex {
				minStartIndex = startIndex
				minPattern := make([]int, len(pattern))
				copy(minPattern, pattern)
				continue
			}
			fmt.Println("-----------------------------")
		}
	}

	if len(minPattern) == 0 {
		return nums, nil, fmt.Errorf("no pattern")
	}

	fmt.Println("response", minPattern, minStartIndex)
	return nums[:minStartIndex], nums[minStartIndex:], nil
}

func main() {
	// fmt.Println(endsIn([]int{1, 2, 3, 4, 2, 3}, []int{1, 2, 3, 4}))
	fmt.Println(CutRepeatingPattern([]int{1, 2, 3, 4, 2, 3}))
}
