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
		} else {
			panic("wrong instruction")
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
	return navigation{graph, instr}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	a, b := integers[0], integers[1]
	result := a * b / GCD(a, b)
	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func all(a []int) bool {
	for i := range a {
		if a[i] == 0 {
			return false
		}
	}
	return true
}

func findSolution(nav navigation) int64 {
	var nodes []string
	for k := range nav.graph {
		if strings.HasSuffix(k, "A") {
			nodes = append(nodes, k)
		}
	}
	cycles := make([]int, len(nodes))
	answer := make([]int, len(nodes))

out:
	for {
		for _, dir := range nav.instructions {
			for i := 0; i < len(nodes); i++ {
				cycles[i]++
				node := nav.graph[nodes[i]][dir]
				if strings.HasSuffix(node, "Z") {
					answer[i] = cycles[i]
					cycles[i] = 0
				}
				nodes[i] = node
			}
			if all(answer) {
				break out
			}
		}
	}

	return int64(LCM(answer...))
}

func main() {
	start := time.Now()
	solution := solvePuzzle("./day8/input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("Solution: ", solution)
}
