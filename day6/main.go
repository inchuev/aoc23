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
	u, err := strconv.ParseUint(s, 10, 64)
	checkErr(err)
	return int64(u)
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

type travel struct {
	time     []int64
	distance []int64
}

func parseInput(lines []string) travel {
	time := fromStrArr(strings.Split(lines[0], ":")[1])
	distance := fromStrArr(strings.Split(lines[1], ":")[1])
	return travel{time, distance}
}

func waysToWin(t, d int64) int64 {
	const accel = 1
	var minTime, maxTime int64
	for i := int64(0); i < t; i++ {
		left := t - i
		traveled := (i * accel) * left
		if traveled > d {
			if minTime == 0 {
				fmt.Println("Found min time to distance:", i)
				minTime = i
			}
		} else {
			if minTime != 0 {
				fmt.Println("Found max time to distance:", i)
				maxTime = i
				break
			}
		}
	}
	return max((maxTime - minTime), 1)
}

func findSolution(data travel) int64 {
	var mul int64 = 1
	for i := 0; i < len(data.distance); i++ {
		t, d := data.time[i], data.distance[i]
		mul *= waysToWin(t, d)
	}
	return mul
}

func main() {
	start := time.Now()
	part1 := solvePuzzle("./day6/input")
	part2 := solvePuzzle("./day6/input2")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution 1 & 2: ", part1, part2)
}
