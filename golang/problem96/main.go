package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sudoku [9][9]int64

func (s *Sudoku) Clone() Sudoku {
	var clone Sudoku
	for i := 0; i < 9; i++ {
		copy(clone[i][:], s[i][:])
	}
	return clone
}

func (s Sudoku) String() string {
	var sb strings.Builder
	for i, row := range s {
		for j, val := range row {
			if val == 0 {
				sb.WriteString(".")
			} else {
				sb.WriteString(fmt.Sprintf("%d", val))
			}
			if j == 2 || j == 5 {
				sb.WriteString(" ")
			}
			if j < 8 {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
		if (i+1)%3 == 0 && i < 8 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (s *Sudoku) IsSolved() bool {
	for i := range 9 {
		for j := range 9 {
			if s[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func (s Sudoku) GetPosibilities(i, j int) []int64 {
	if s[i][j] != 0 {
		return []int64{s[i][j]}
	}

	posibilities := map[int64]struct{}{}
	for i := int64(1); i < 10; i++ {
		posibilities[i] = struct{}{}
	}

	// Remove row nums
	for row := range 9 {
		v := s[row][j]
		if v != 0 {
			delete(posibilities, v)
		}
	}

	// Remove row columns
	for col := range 9 {
		v := s[i][col]
		if v != 0 {
			delete(posibilities, v)
		}
	}

	// Remove numbers in the 3x3 cell
	boxRowStart := (i / 3) * 3
	boxColStart := (j / 3) * 3
	for row := boxRowStart; row < boxRowStart+3; row++ {
		for col := boxColStart; col < boxColStart+3; col++ {
			v := s[row][col]
			if v != 0 {
				delete(posibilities, v)
			}
		}
	}

	var result []int64
	for k := int64(1); k < 10; k++ {
		if _, ok := posibilities[k]; ok {
			result = append(result, k)
		}
	}

	return result
}

func (s Sudoku) SatisfyConstraints() Sudoku {
outerLoop:
	for {
		c := int64(0)
		for i := range 9 {
			for j := range 9 {
				if s[i][j] == 0 {
					posibilities := s.GetPosibilities(i, j)
					if len(posibilities) == 1 {
						s[i][j] = posibilities[0]
						c++
					}
				}
			}
		}
		if c == 0 {
			break outerLoop
		}
	}
	return s
}

func (s Sudoku) Solve() (Sudoku, bool) {
	if s.IsSolved() {
		return s, true
	}

	s = s.SatisfyConstraints()

	if s.IsSolved() {
		return s, true
	}

	var row, col int
	found := false
outerLoop:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[i][j] == 0 {
				row, col = i, j
				found = true
				break outerLoop
			}
		}
	}

	if !found {
		return s, true
	}

	possibilities := s.GetPosibilities(row, col)

	for _, val := range possibilities {
		newS := s.Clone()
		newS[row][col] = val

		solution, solved := newS.Solve()
		if solved {
			return solution, solved
		}

	}

	return s, s.IsSolved()
}

func ParseGame(input []string) Sudoku {
	game := Sudoku{}
	for i := range 9 {
		for j := range 9 {
			val, err := strconv.ParseInt(string(input[i][j]), 10, 64)
			if err != nil {
				panic(err)
			}

			game[i][j] = val
		}
	}

	return game
}

func main() {
	start := time.Now()
	games := map[string]Sudoku{}
	f, err := os.Open("./p096_sudoku.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	lineCount := 0
	var gameName string
	gameStr := []string{}
	for scanner.Scan() {
		if lineCount == 0 {
			gameName = strings.TrimSpace(scanner.Text())
		} else {
			gameStr = append(gameStr, strings.TrimSpace(scanner.Text()))
		}

		if lineCount == 9 {
			game := ParseGame(gameStr)
			games[gameName] = game
			gameStr = []string{}
			lineCount = 0
			continue
		}

		lineCount += 1
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	solutions := map[string]Sudoku{}
	for gameName, game := range games {
		solution, solved := game.Solve()
		if !solved {
			panic(fmt.Sprintf("game %v not solved!", gameName))
		}
		solutions[gameName] = solution
	}

	sum := int64(0)
	for _, solution := range solutions {
		sum += solution[0][0]*100 + solution[0][1]*10 + solution[0][2]
	}

	template := "Sum of 3-digit numbers at top left = %v (elapsed = %v)\n"
	elapsed := time.Since(start)
	fmt.Printf(template, sum, elapsed)
}
