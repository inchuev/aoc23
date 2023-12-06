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

func fromStrArr(arr []string) []int64 {
	var uarr []int64
	for i := range arr {
		if arr[i] == "" {
			continue
		}
		uarr = append(uarr, fromStr(arr[i]))
	}
	return uarr
}

func solvePuzzle(path string) (sum int64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	seeds := getSeeds(lines[0])
	rangeMap := parseInput(lines)
	low := findLocation(rangeMap, seeds)

	return low
}

type mapping struct {
	src int64
	dst int64
	rng int64
}

func parseInput(lines []string) [][]mapping {
	mappings := make([][]mapping, 7)
	cat := -1
	for i := 1; i < len(lines); i++ {
		l := lines[i]
		if l == "" {
			continue
		}
		if strings.Contains(l, "map:") {
			cat++
			continue
		}
		arr := fromStrArr(strings.Split(l, " "))
		mappings[cat] = append(mappings[cat], mapping{arr[0], arr[1], arr[2]})
	}
	return mappings
}

func getSeeds(l string) []int64 {
	return fromStrArr(strings.Split(strings.Split(l, ":")[1], " "))
}

func overlap(val int64, m mapping) (int64, bool) {
	if m.dst > val {
		return 0, false
	}
	if m.dst == val {
		return m.src, true
	}
	if val <= (m.dst + m.rng) {
		if (val - m.dst) > 0 {
			return val - m.dst + m.src, true
		} else {
			return val + m.src, true
		}
	}
	return 0, false
}

func overlapInCategory(cat []mapping, val int64) int64 {
	for _, m := range cat {
		if res, ok := overlap(val, m); ok {
			return res
		}
	}
	return val
}

func findLocation(rm [][]mapping, seeds []int64) int64 {
	var loc, val int64 = int64(^uint64(0) >> 1), 0
	for _, seed := range seeds {
		val = seed
		for _, cat := range rm {
			val = overlapInCategory(cat, val)
		}
		if val < loc {
			loc = val
		}
	}
	return loc
}

func main() {
	start := time.Now()
	res := solvePuzzle("./day5/input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", res)
}
