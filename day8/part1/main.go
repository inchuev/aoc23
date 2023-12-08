package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fromStr(s []string) []int {
	directions := make([]int, len(s))
	for i, dir := range s {
		if dir == "L" {
			directions[i] = 0
		} else if dir == "R" {
			directions[i] = 1
		}
	}
	return directions
}

func solvePuzzle(path string) (sum int64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	parsed := parseInput(lines)
	result := findSolution(parsed)

	return result
}

type navigation struct {
	graph        map[string][]string
	instructions []int
}

func parseInput(lines []string) navigation {
	graph := map[string][]string{}
	instr := fromStr(strings.Split(strings.TrimSpace(lines[0]), ""))
	r := strings.NewReplacer(
		"(", "",
		")", "",
		" ", "",
	)
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		s := strings.Split(r.Replace(lines[i]), "=")
		node := s[0]
		conns := strings.Split(s[1], ",")
		graph[node] = conns
	}
	fmt.Println(graph)
	return navigation{graph, instr}
}

func findSolution(nav navigation) int64 {
	var steps int64
	node := "AAA"

	for {
		for _, dir := range nav.instructions {
			steps++
			node = nav.graph[node][dir]
			if node == "ZZZ" {
				return steps
			}
		}
	}
}

func main() {
	start := time.Now()
	// Part 1 is identical, but without special joker processing.
	part2 := solvePuzzle("./day8/input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", part2)
}
