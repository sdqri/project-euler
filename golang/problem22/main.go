package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func score(s string) uint {
	sum := 0
	r := 0
	for _, c := range strings.ToUpper(s) {
		sum += int(c)
		r++
	}
	return uint(sum - r*64)
}

func main() {
	file, err := os.Open("./0022_names.txt")
	if err != nil {
		panic(err)
	}

	fileStats, err := file.Stat()
	if err != nil {
		panic(err)
	}
	data := make([]byte, fileStats.Size())
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}
	names := strings.Split(strings.ReplaceAll(string(data), "\"", ""), ",")

	sort.Sort(sort.StringSlice(names))

	var sum uint = 0
	for i, name := range names {
		sum += uint(i+1) * score(name)
	}
	fmt.Printf("score of file = %v\n", sum)
}
