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

func solvePuzzle(path string) (sum uint64) {
	b, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(string(b), "\n")
	for i := range lines {
		if len(lines[i]) > 0 {
			sum += findNumber(lines[i])
		}
	}
	return sum
}

func getDigit(r rune) uint64 {
	if r == '1' {
		return 1
	} else if r == '2' {
		return 2
	} else if r == '3' {
		return 3
	} else if r == '4' {
		return 4
	} else if r == '5' {
		return 5
	} else if r == '6' {
		return 6
	} else if r == '7' {
		return 7
	} else if r == '8' {
		return 8
	} else if r == '9' {
		return 9
	} else {
		return 0
	}
}

var R = strings.NewReplacer(
	"oneight", "18",
	"threeight", "38",
	"twone", "21",
	"fiveight", "58",
	"eightwo", "82",
	"eighthree", "83",
	"nineight", "98",
	"one", "1",
	"two", "2",
	"three", "3",
	"four", "4",
	"five", "5",
	"six", "6",
	"seven", "7",
	"eight", "8",
	"nine", "9",
)

func findNumber(s string) uint64 {
	s = R.Replace(s)
	res := _findNumber([]rune(s))
	return res
}

// Find two digits in the string and glue them together.
func _findNumber(s []rune) uint64 {
	var f, l uint64
	i, j := 0, len(s)-1
	for f == 0 || l == 0 {
		if f == 0 {
			f = getDigit(s[i])
		}
		if l == 0 {
			l = getDigit(s[j])
		}
		i++
		j--
	}
	return f*10 + l
}

func main() {
	start := time.Now()
	res := solvePuzzle("./input")
	end := time.Since(start)
	fmt.Println("Took: ", end)
	fmt.Println("total: ", res)
}
