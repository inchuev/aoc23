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

func fromStr(s string) int64 {
	u, err := strconv.ParseInt(s, 10, 64)
	checkErr(err)
	return u
}

func fromStrArr(s string) []int64 {
	strArr := strings.Split(strings.TrimSpace(s), " ")
	iarr := make([]int64, 0, len(strArr))
	for i := range strArr {
		if strArr[i] == "" {
			continue
		}
		iarr = append(iarr, fromStr(strArr[i]))
	}
	return iarr
}

func solvePuzzle(path string) (sum int64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	parsed := parseInput(lines)
	result := findSolution(parsed)

	return result
}

func parseInput(lines []string) [][]int64 {
	arr := make([][]int64, len(lines))
	for i := range lines {
		arr[i] = fromStrArr(lines[i])
	}
	return arr
}

func findSolution(sequences [][]int64) int64 {
	var sum int64
	for i := range sequences {
		seq := sequences[i]
		sum += nth(seq)
	}
	return sum
}

func nth(seq []int64) int64 {
	var diffs []int64
	s := seq
	for {
		diff := make([]int64, len(s)-1)
		var last int64
		for i := 0; i < len(s)-1; i++ {
			diff[i] = s[i+1] - s[i]
			last = diff[i]
		}
		s = diff
		diffs = append(diffs, last)
		if last == 0 {
			break
		}
	}
	x := sum(diffs)

	return seq[len(seq)-1] + x
}

func sum(a []int64) int64 {
	var s int64
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	start := time.Now()
	part2 := solvePuzzle("./day9/input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", part2)
}
