package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "embed"
)

//go:embed 0059_cipher.txt
var text string

func StretchKey(key []byte, length int) []byte {
	if length <= 0 {
		return nil
	}

	keyLength := len(key)
	repeats := length / keyLength
	remaining := length % keyLength
	return append(bytes.Repeat(key, repeats), key[:remaining]...)
}

func xorBytes(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("different length xor")
	}
	result := make([]byte, len(a))
	for i := range a {
		result[i] = a[i] ^ b[i]
	}
	return result
}

func hasCommonWord(source []byte) bool {
	sourceText := string(source)
	if strings.Contains(sourceText, " the ") {
		return true
	}

	return false
}

func main() {
	r := strings.NewReader(text)
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var byteArray []byte
	for _, word := range records[0] {
		num, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}

		byteArray = append(byteArray, byte(num))
	}

	var key []byte
	var decodedBytes []byte
	start := time.Now()
outerLoop:
	for i := byte(97); i < 123; i++ {
		for j := byte(97); j < 123; j++ {
			for k := byte(97); k < 123; k++ {
				key = []byte{i, j, k}
				stretchedKey := StretchKey(key, len(byteArray))
				decodedBytes = xorBytes(stretchedKey, byteArray)
				if hasCommonWord(decodedBytes) {
					break outerLoop
				}
			}
		}
	}

	sum := int64(0)
	for _, b := range decodedBytes {
		sum += int64(b)
	}

	elapsed := time.Since(start)
	fmt.Printf("found key = %s (elapsed = %v)\nsum of chars = %d\ntext = %s \n", string(key), elapsed, sum, string(decodedBytes))
}
