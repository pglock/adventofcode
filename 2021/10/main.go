package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lines = append(lines, lineStr)
	}
	return lines
}

func isOpening(char rune) bool {
	return strings.ContainsRune("{[(<", char)
}

func isMatching(opening string, closing string) bool {
	check := make(map[string]string)
	check["("] = ")"
	check["["] = "]"
	check["{"] = "}"
	check["<"] = ">"
	return check[opening] == closing
}

func getScore(char string) int {
	scores := make(map[string]int)
	scores[")"] = 3
	scores["]"] = 57
	scores["}"] = 1197
	scores[">"] = 25137
	return scores[char]
}

func processLine(line string) int {
	stack := make([]string, 0)
	top := "c"
	for _, c := range line {
		asString := string(c)
		if isOpening(c) {
			stack = append(stack, asString)
		} else {
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if !isMatching(top, asString) {
				return getScore(asString)
			}
		}
	}
	return 0
}

func main() {
	dat := readData("input.txt")
	total := 0
	for _, line := range dat {
		total += processLine(line)
	}
	fmt.Println(total)
	fmt.Println("finished")
}
