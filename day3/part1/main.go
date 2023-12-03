package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func solvePuzzle(path string) (sum uint64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	for i := range lines {
		if len(lines[i]) > 0 {
			sum += parts(lines, i)
		}
	}
	return sum
}

func isDigit(c rune) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func parts(rows []string, i int) (sum uint64) {
	var pos, startPos, endPos int = 0, -1, -1
	var char rune
	for pos, char = range rows[i] {
		if startPos == -1 && isDigit(char) {
			startPos = pos
		}
		if startPos != -1 && !isDigit(char) {
			endPos = pos - 1
		}
		if startPos != -1 && endPos != -1 {
			sum += checkAdjacent(rows, i, startPos, endPos)
			startPos, endPos = -1, -1
		}
	}
	if startPos != -1 {
		endPos = pos
		sum += checkAdjacent(rows, i, startPos, endPos)
	}

	return sum
}

func checkAdjacent(rows []string, rowN, fst, lst int) uint64 {
	num := func() uint64 {
		s := rows[rowN][fst : lst+1]
		res, err := strconv.ParseUint(s, 10, 64)
		checkErr(err)
		return res
	}
	// curr row: pos - 1, pos + 1
	// prev row: pos, pos - 1, pos + 1
	// next row: pos, pos - 1, pos + 1
	// all rows are the same width
	rowLen := len(rows[rowN])
	columns := []int{fst, lst}
	if (fst - 1) >= 0 {
		columns = append(columns, fst-1)
	}
	if (fst + 1) < rowLen {
		columns = append(columns, fst+1)
	}
	if (lst - 1) >= 0 {
		columns = append(columns, lst-1)
	}
	if (lst + 1) < rowLen {
		columns = append(columns, lst+1)
	}
	checkRowsN := []int{rowN}
	if (rowN - 1) >= 0 {
		checkRowsN = append(checkRowsN, rowN-1)
	}
	if (rowN + 1) < len(rows) {
		checkRowsN = append(checkRowsN, rowN+1)
	}
	for _, rn := range checkRowsN {
		row := []rune(rows[rn])
		for _, col := range columns {
			if isSymbol(row[col]) {
				return num()
			}
		}
	}

	return 0
}

func isSymbol(r rune) bool {
	return r != '.' && !isDigit(r)
}

func main() {
	start := time.Now()
	res := solvePuzzle("./input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", res)
}
