package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
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

func fromHex(s string) int64 {
	u, err := strconv.ParseUint(s, 16, 64)
	checkErr(err)
	return int64(u)
}

func solvePuzzle(path string) (sum int64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	parsed := parseInput(lines)
	result := findSolution(parsed)

	return result
}

type record struct {
	hand int64
	bid  int64
}

type STRENGTH int

const (
	HIGH_CARD STRENGTH = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_KIND
	FULL_HOUSE
	FOUR_OF_KIND
	FIVE_OF_KIND
)

func classify(hand []rune) STRENGTH {
	m := make(map[rune]int, len(hand))
	joker := 0
	for _, r := range hand {
		if r == 'J' {
			joker++
		}
		m[r]++
	}
	pairs := 0
	threes := 0
	typ := HIGH_CARD
	for k := range m {
		switch m[k] {
		case 5:
			typ = FIVE_OF_KIND
		case 4:
			typ = FOUR_OF_KIND
		case 3:
			threes++
		case 2:
			pairs++
		case 1:
		default:
			panic("should never happen")
		}
	}
	if threes == 1 && pairs == 1 {
		typ = FULL_HOUSE
	} else if threes == 1 {
		typ = THREE_OF_KIND
	} else if pairs == 2 {
		typ = TWO_PAIR
	} else if pairs == 1 {
		typ = ONE_PAIR
	}
	typ = reclassify(typ, joker)
	return typ
}

func reclassify(s STRENGTH, joker int) STRENGTH {
	switch s {
	case FOUR_OF_KIND:
		switch joker {
		case 1, 4:
			return FIVE_OF_KIND
		}
	case FULL_HOUSE:
		switch joker {
		case 1:
			return FOUR_OF_KIND
		case 2, 3:
			return FIVE_OF_KIND
		}
	case THREE_OF_KIND:
		switch joker {
		case 1, 3:
			return FOUR_OF_KIND
		}
	case TWO_PAIR:
		switch joker {
		case 1:
			return FULL_HOUSE
		case 2:
			return FOUR_OF_KIND
		}
	case ONE_PAIR:
		switch joker {
		case 1, 2:
			return THREE_OF_KIND
		}
	case HIGH_CARD:
		switch joker {
		case 1:
			return ONE_PAIR
		}
	}
	return s
}

func parseInput(lines []string) map[STRENGTH][]record {
	r := strings.NewReplacer(
		"A", "F",
		"K", "E",
		"Q", "D",
		"J", "1",
		"T", "B",
	)
	buckets := make(map[STRENGTH][]record)
	for i := range lines {
		l := strings.Split(lines[i], " ")
		hand, bid := []rune(l[0]), fromStr(l[1])
		typ := classify(hand)
		buckets[typ] = append(buckets[typ], record{fromHex(r.Replace(l[0])), bid})
	}
	return buckets
}

func findSolution(buckets map[STRENGTH][]record) int64 {
	var winnings, currentRank int64
	for _, bn := range []STRENGTH{HIGH_CARD, ONE_PAIR, TWO_PAIR, THREE_OF_KIND, FULL_HOUSE, FOUR_OF_KIND, FIVE_OF_KIND} {
		bkt := buckets[bn]
		slices.SortFunc(bkt, func(a, b record) int {
			return cmp.Compare(a.hand, b.hand)
		})
		for _, rec := range bkt {
			currentRank++
			winnings += rec.bid * currentRank
		}
	}

	return winnings
}

func main() {
	start := time.Now()
	// Part 1 is identical, but without special joker processing.
	part2 := solvePuzzle("./day7/input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", part2)
}
