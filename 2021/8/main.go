package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type display struct {
	patterns []string
	samples  []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(filename string) []display {
	file, err := os.Open(filename)
	check(err)
	lines := make([]display, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		splits := strings.Split(lineStr, "|")
		lines = append(lines, display{strings.Fields(splits[0]), strings.Fields(splits[1])})
	}
	return lines
}

func countUniques(lines []display) int {
	cnt := 0
	for _, line := range lines {
		for _, sample := range line.samples {
			if len(sample) == 2 || len(sample) == 3 || len(sample) == 4 || len(sample) == 7 {
				cnt += 1
			}
		}
	}
	return cnt
}

func main() {
	dat := readData("input.txt")
	fmt.Println(countUniques(dat))
	fmt.Println("finished")
}
