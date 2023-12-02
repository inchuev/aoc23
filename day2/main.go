package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var GRE = regexp.MustCompile(`^Game (\d+)?: (.*)$`)
var BRE = regexp.MustCompile(`(?U)^(\d+)? (green|red|blue)?$`)
var L = map[string]uint64{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func solvePuzzle(path string) (sum uint64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(string(b), "\n")
	for i := range lines {
		if len(lines[i]) > 0 {
			//sum += gameIDPossible(lines[i])
			sum += ballsGame(lines[i])
		}
	}
	return sum
}

func ballsGame(l string) uint64 {
	m := GRE.FindStringSubmatch(l)
	var r, r1, g, g1, b, b1 uint64
	sets := strings.Split(m[2], ";")
	for _, set := range sets {
		r1, g1, b1 = balls(set)
		r, g, b = max(r, r1), max(g, g1), max(b, b1)
	}
	return r * g * b
}

func balls(s string) (r uint64, g uint64, b uint64) {
	balls := strings.Split(s, ",")
	for _, ballStr := range balls {
		balls := BRE.FindStringSubmatch(strings.TrimSpace(ballStr))
		count, err := strconv.ParseUint(balls[1], 10, 64)
		checkErr(err)
		color := balls[2]
		switch color {
		case "red":
			r = count
		case "green":
			g = count
		case "blue":
			b = count
		default:
			panic("unknown color!")
		}
	}
	return r, g, b
}

// Returns game ID if the game is possible, 0 otherwise.
func gameIDPossible(l string) uint64 {
	m := GRE.FindStringSubmatch(l)
	gameID, err := strconv.ParseUint(m[1], 10, 64)
	checkErr(err)
	sets := strings.Split(m[2], ";")
	for _, set := range sets {
		if !setPossible(set) {
			return 0
		}
	}
	return gameID
}

func setPossible(s string) bool {
	balls := strings.Split(s, ",")
	for _, ballStr := range balls {
		balls := BRE.FindStringSubmatch(strings.TrimSpace(ballStr))
		count, err := strconv.ParseUint(balls[1], 10, 64)
		checkErr(err)
		color := balls[2]
		if allowed := L[color]; allowed < count {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	res := solvePuzzle("./input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", res)
}
