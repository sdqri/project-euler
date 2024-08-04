package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func copyMap[K comparable, V any](original map[K]V) map[K]V {
	copied := make(map[K]V, len(original))

	for key, value := range original {
		copied[key] = value
	}

	return copied
}

func isValidTriangleNumber(x int64) bool {
	// n^2 + n - 2x = 0
	// discriminant = b^2 - 4ac = 1 - 4 * 1 * -2x
	discriminant := 1 + 8*x
	root := int64(math.Sqrt(float64(discriminant)))
	return root*root == discriminant
}

func isValidSquareNumber(x int64) bool {
	if x < 0 {
		return false // Negative numbers cannot be perfect squares
	}
	root := int64(math.Sqrt(float64(x)))
	return root*root == x
}

func isValidPentagonalNumber(x int64) bool {
	// 3n^2 - n - 2x = 0
	// discriminant = b^2 - 4ac = 1 - 4 * 3 * -2x
	discriminant := 1 + 24*x
	root := int64(math.Sqrt(float64(discriminant)))
	return root*root == discriminant && (1+root)%6 == 0
}

func isValidHexagonalNumber(x int64) bool {
	// 2n^2 - n - x = 0
	// discriminant = b^2 - 4ac = 1 - 4 * 2 * -1x
	discriminant := 1 + 8*x
	root := int64(math.Sqrt(float64(discriminant)))
	return root*root == discriminant && (1+root)%4 == 0
}

func isValidHeptagonalNumber(x int64) bool {
	// 5n^2 - 3n - 2x = 0
	// discriminant = b^2 - 4ac = 9 + 4 * 5 * (-2x)
	discriminant := 9 + 40*x
	root := int64(math.Sqrt(float64(discriminant)))
	return root*root == discriminant && (3+root)%10 == 0
}

func isValidOctagonalNumber(x int64) bool {
	// 3n^2 - 2n - x = 0
	// discriminant = b^2 - 4ac = 4 - 4 * 3 * (-x)
	discriminant := 4 + 12*x
	root := int64(math.Sqrt(float64(discriminant)))
	return root*root == discriminant && (2+root)%6 == 0
}

var prefixes = []struct {
	isValidFunc func(int64) bool
	prefix      string
}{
	{isValidTriangleNumber, "tri"},
	{isValidSquareNumber, "sqr"},
	{isValidPentagonalNumber, "pen"},
	{isValidHexagonalNumber, "hex"},
	{isValidHeptagonalNumber, "hep"},
	{isValidOctagonalNumber, "oct"},
}

func isSetCyclic(cyclicDigits int, elems ...int) bool {
	if len(elems) == 0 {
		return false
	}

	strElems := make([]string, len(elems))
	for i, elem := range elems {
		strElem := strconv.Itoa(elem)
		elemLen := len(strElem)
		if elemLen < cyclicDigits {
			return false
		}
		strElems[i] = strElem
	}

	if len(elems) > 1 {
		for i := 0; i < len(elems)-1; i++ {
			xlastDigits := strElems[i][len(strElems[i])-cyclicDigits:]
			yfirstDigits := strElems[i+1][:cyclicDigits]
			if xlastDigits != yfirstDigits {
				return false
			}
		}
	}

	lastIndex := len(elems) - 1
	firstElem := strElems[0]
	lastElem := strElems[lastIndex]
	lastElemEnd := lastElem[len(lastElem)-cyclicDigits:]
	firstElemStart := firstElem[:cyclicDigits]
	if firstElemStart != lastElemEnd {
		return false
	}

	return true
}

func rotateToSmallest(list []string) []string {
	if len(list) == 0 {
		return list
	}

	minIndex := 0
	for i, v := range list {
		if v < list[minIndex] {
			minIndex = i
		}
	}

	return append(list[minIndex:], list[:minIndex]...)
}

func getPathKey(path []string) string {
	return fmt.Sprintf("%v", path)
}

func findCyclesWithLength(length int, graph map[string][]string) [][]string {
	var allCycles [][]string

	var dfs func(path []string, visited map[string]bool)
	dfs = func(path []string, visited map[string]bool) {
		start := path[0]
		head := path[len(path)-1]
		edges := graph[head]
		if len(path) == length {
			if contains(edges, start) {
				cycle := make([]string, len(path))
				copy(cycle, path)
				allCycles = append(allCycles, cycle)
			}
			return
		}

		for _, edge := range edges {
			if !visited[edge] {
				newPath := append(path, edge)
				new_visited := copyMap(visited)
				new_visited[edge] = true
				dfs(newPath, new_visited)
			}
		}
	}

	for node := range graph {
		if strings.HasPrefix(node, "tri") {
			dfs([]string{node}, map[string]bool{node: true})
		}
	}

	uniqueCycles := map[string][]string{}
	for _, cycle := range allCycles {
		smallestCycle := rotateToSmallest(cycle)
		uniqueCycles[getPathKey(smallestCycle)] = smallestCycle
	}

	uniqueCyclesSlice := [][]string{}
	for _, cycle := range uniqueCycles {
		uniqueCyclesSlice = append(uniqueCyclesSlice, cycle)
	}

	return uniqueCyclesSlice
}

func main() {
	start := time.Now()
	cyclicDigits := 2
	cdMap := map[string][]string{}
	elemsSet := map[string]struct{}{}

	for i := int64(1e3); i < int64(1e4); i++ {
		iStr := strconv.Itoa(int(i))
		icd := iStr[:cyclicDigits] // i's cyclic digits
		v, ok := cdMap[icd]

		values := []string{}
		for _, p := range prefixes {
			if p.isValidFunc(i) {
				value := fmt.Sprintf("%s%s", p.prefix, iStr)
				values = append(values, value)
			}
		}

		if len(values) == 0 {
			continue
		}

		if !ok {
			cdMap[icd] = []string{}
		}
		v = append(v, values...)

		cdMap[icd] = v
		for _, value := range values {
			elemsSet[value] = struct{}{}
		}
	}

	graph := map[string][]string{}
	for elem := range elemsSet {
		startIdx := len(elem) - cyclicDigits
		elemNum := elem[startIdx:]
		cd := elemNum[:cyclicDigits]
		value, ok := cdMap[cd]
		if ok {
			graph[elem] = value
		} else {
			graph[elem] = []string{}
		}
	}

	cycles := findCyclesWithLength(6, graph)

	setOfSix := [][]string{}
	for _, c := range cycles {
		polygonalSet := map[string]struct{}{}
		for _, elem := range c {
			polygonalSet[elem[:3]] = struct{}{}
		}
		if len(polygonalSet) >= 6 {
			setOfSix = append(setOfSix, c)
		}
	}

	elapsed := time.Since(start)
	sums := []int{}
	for _, s := range setOfSix {
		sum := 0
		for _, e := range s {
			v, err := strconv.Atoi(e[3:])
			if err != nil {
				panic(err)
			}
			sum += v
		}
		sums = append(sums, sum)
	}

	fmt.Printf("setOfSix = %v \n", setOfSix)
	fmt.Printf("sum of setOfSix = %v (elapsed = %v)\n", sums, elapsed)
}
