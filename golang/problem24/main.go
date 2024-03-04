package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func removeElement[T any](slice []T, index int) []T {
	result := make([]T, 0)
	for i, element := range slice {
		if i != index {
			result = append(result, element)
		}
	}
	return result
}

func sliceDifference[T comparable](slice1, slice2 []T) []T {
	difference := []T{}

	lookup := make(map[T]bool)
	for _, item := range slice2 {
		lookup[item] = true
	}

	for _, item := range slice1 {
		if !lookup[item] {
			difference = append(difference, item)
		}
	}
	return difference
}

func factorial(x int) int {
	result := 1
	for i := range x {
		result *= i + 1
	}
	return result
}

func FindNthLexicographicPermutations(alphabet []string, n int) string {
	permuations, remainedN := findNthLexicographicPermutations(alphabet, n)
	remainedAlphabet := sliceDifference(alphabet, permuations)
	if len(remainedAlphabet) == 1 {
		permuations = append(permuations, remainedAlphabet[0])
		return strings.Join(permuations, "")
	}

	remainedPermutations := Permute(remainedAlphabet)
	remainedPermutationsStringSlice := make([]string, 0)
	for _, permutation := range remainedPermutations {
		remainedPermutationsStringSlice = append(remainedPermutationsStringSlice, strings.Join(permutation, ""))
	}
	sort.Sort(sort.StringSlice(remainedPermutationsStringSlice))
	remainedPerm := remainedPermutationsStringSlice[remainedN-1]
	for _, letter := range remainedPerm {
		permuations = append(permuations, string(letter))
	}

	return strings.Join(permuations, "")
}

func findNthLexicographicPermutations(alphabet []string, n int) ([]string, int) {
	permutation := make([]string, 0)
	sort.Sort(sort.StringSlice(alphabet))
outerLoop:
	for len(alphabet) > 0 {
		selectedIndex := -1
		states := factorial(len(alphabet) - 1)
		for i := range alphabet {
			if ((i + 1) * states) < n {
				selectedIndex = i
			} else if selectedIndex != -1 { //if nothing is selected and n is smaller
				break
			} else { // if sth is selected and n is smaller
				break outerLoop
			}
		}
		n -= (selectedIndex + 1) * states
		permutation = append(permutation, alphabet[selectedIndex+1])
		alphabet = removeElement(alphabet, selectedIndex+1)
	}

	return permutation, n
}

func Permute[T any](alphabet []T) [][]T {
	permutations := make([][]T, 0)
	permute(alphabet, []T{}, &permutations)
	return permutations
}

func permute[T any](alphabet, current []T, result *[][]T) {
	if len(alphabet) == 0 {
		*result = append(*result, current)
		return
	}

	for i := range alphabet {
		newAlphabet := make([]T, len(alphabet)-1)
		copy(newAlphabet[:i], alphabet[:i])
		copy(newAlphabet[i:], alphabet[i+1:])
		newCurrent := append(current, alphabet[i])
		permute(newAlphabet, newCurrent, result)
	}
}

func main() {
	alphabet := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	n := 1000000

	start := time.Now()
	permutations := Permute(alphabet)
	permutationsStringSlice := make([]string, 0)
	for _, permutation := range permutations {
		permutationsStringSlice = append(permutationsStringSlice, strings.Join(permutation, ""))
	}
	sort.Sort(sort.StringSlice(permutationsStringSlice))
	result := permutationsStringSlice[n-1]
	elapsedBruteforce := time.Since(start)
	start = time.Now()
	newResult := FindNthLexicographicPermutations(alphabet, n)
	elapsedNewMethod := time.Since(start)
	fmt.Printf(
		"brute-force method: %dth lexicographic permutations = %s (elapsed time = %v)\n",
		n,
		result,
		elapsedBruteforce,
	)
	fmt.Printf(
		"new method: %dth lexicographic permutations = %s (elapsed time = %v)\n",
		n,
		newResult,
		elapsedNewMethod,
	)
}
