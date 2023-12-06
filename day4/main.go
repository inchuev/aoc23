package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var CARDS = make([]uint64, 202)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fromStr(s string) uint64 {
	u, err := strconv.ParseUint(s, 10, 64)
	checkErr(err)
	return u
}

func fromStrArr(arr []string) map[uint64]struct{} {
	uarr := make(map[uint64]struct{}, len(arr))
	for i := range arr {
		if arr[i] == "" {
			continue
		}
		uarr[fromStr(arr[i])] = struct{}{}
	}
	return uarr
}

func solvePuzzle(path string) (sum uint64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	for i := range lines {
		if len(lines[i]) > 0 {
			CARDS[i] += 1
			// part 1:
			//sum += scratchCardValue(lines[i])
			processScratchCard(lines, i)
		}
	}
	sum += totalScrachCards()

	return sum
}

func winning(win, rest map[uint64]struct{}) []uint64 {
	var w []uint64
	for k := range win {
		if _, ok := rest[k]; ok {
			w = append(w, k)
		}
	}
	return w
}

func totalScrachCards() uint64 {
	var sum uint64
	for i := range CARDS {
		sum += CARDS[i]
	}
	return sum
}

func processScratchCard(lines []string, card int) {
	l := lines[card]
	arr := strings.Split(strings.Split(l, ":")[1], "|")
	win, rest := fromStrArr(strings.Split(strings.TrimSpace(arr[0]), " ")), fromStrArr(strings.Split(strings.TrimSpace(arr[1]), " "))
	winNums := winning(win, rest)
	wl := len(winNums)
	for i := card + 1; i <= card+wl; i++ {
		if i == len(CARDS) {
			break
		}
		CARDS[i] += CARDS[card]
		/*
			// Recursive solution is more natural, but the performance is horrible. (~16 seconds on my laptop)
			CARDS[i]++
			processScratchCard(lines, i)
		*/
	}
}

func scratchCardValue(l string) uint64 {
	arr := strings.Split(strings.Split(l, ":")[1], "|")
	win, rest := fromStrArr(strings.Split(strings.TrimSpace(arr[0]), " ")), fromStrArr(strings.Split(strings.TrimSpace(arr[1]), " "))
	winNums := winning(win, rest)
	wl := len(winNums)
	switch wl {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return 1 << (len(winNums) - 1)
	}
}

func main() {
	start := time.Now()
	res := solvePuzzle("./day4/input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", res)
}
