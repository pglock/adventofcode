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
	numbers [10]string
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
		lines = append(lines, display{patterns: strings.Fields(splits[0]), samples: strings.Fields(splits[1])})
	}
	lines = setUniques(lines)
	for i:=0; i<len(lines); i++ {
		setLen6(&lines[i])
		setLen5(&lines[i])
	}
	return lines
}

func filterLength(patterns []string, n int) []string {
	p := make([]string, 0)
	for _, pattern := range patterns {
		if len(pattern) == n {
			p = append(p, pattern)
		}
	}
	return p
}

func hasAll(check string, subset string) bool {
	for _, c := range subset {
		if !strings.ContainsRune(check, c) {
			return false
		}
	}
	return true
}

func substract(left string, right string) string {
	result := make([]rune, 0)
	for _, c := range left {
		if !strings.ContainsRune(right, c){
			result = append(result, c)
		}
	}
	return string(result)
}

func setLen6(display *display){
	patterns := filterLength(display.patterns, 6)
	for _, pattern := range patterns {
		// set pattern for 9
		if hasAll(pattern, display.numbers[4]) {
			display.numbers[9] = pattern
		} else if len(substract(pattern, display.numbers[1])) == 5 {
		// set pattern for 6
			display.numbers[6] = pattern
		} else {
		// set pattern for 0
			display.numbers[0] = pattern
		}
	}
}

func setLen5(display *display){
	patterns := filterLength(display.patterns, 5)
	for _, pattern := range patterns {
		// 6 - 5 is one char left
		if len(substract(display.numbers[6], pattern)) == 1 {
			display.numbers[5] = pattern
		} else if len(substract(pattern, display.numbers[1])) == 4 {
			// 2 - 1 is only one character less
			display.numbers[2] = pattern
		} else {
			display.numbers[3] = pattern
		}
	}
}


func setUniques(lines []display) []display {
	for i:=0; i<len(lines);i++ {
		for _, sample := range lines[i].patterns {
			if len(sample) == 2 {
				lines[i].numbers[1] = sample
			}
			if len(sample) == 3 {
				lines[i].numbers[7] = sample
			}
			if len(sample) == 4 {
				lines[i].numbers[4] = sample
			}
			if len(sample) == 7 {
				lines[i].numbers[8] = sample
			}
		}
	}
	return lines
}

func sumSamples(displays []display) int {
	sum := 0
	for _, display := range displays {
		sum += getNumber(display)
	}
	return sum
}

func getNumber(display display) int {
	multiplier := []int{1000, 100, 10, 1}
	number := 0
	for i:=0; i<4; i++{
		number += multiplier[i] * sampleNumber(display, i)
	}
	return number
}

func equals(left string, right string) bool {
	if hasAll(left, right) && len(left) == len(right) {
		return true
	}
	return false
}

func sampleNumber(display display, i int) int {
	res := -1
	for j, pattern := range display.numbers {
		if equals(pattern, display.samples[i]) {
			res = j
		}
	}
	return res
}

func main() {
	dat := readData("input.txt")
	fmt.Println(sumSamples(dat))
	fmt.Println("finished")
}
