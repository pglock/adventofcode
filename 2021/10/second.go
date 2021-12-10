package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
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
	scores["("] = 1
	scores["["] = 2
	scores["{"] = 3
	scores["<"] = 4
	return scores[char]
}

func lineState(line string) int {
	stack := make([]string, 0)
	top := "c"
	for _, c := range line {
		asString := string(c)
		if isOpening(c) {
			stack = append(stack, asString)
		} else {
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if !isMatching(top, asString) {
				return -1
			}
		}
	}
	if len(stack) > 0 {
		return 1
	} else {
		return 0
	}
}

func filterIncomplete(lines []string) []string {
	incompletes := make([]string, 0)
	for _, line := range lines {
		if lineState(line) ==  1 {
			incompletes = append(incompletes, line)
		}
	}
	return incompletes
}

func getStack(line string) []string {

	stack := make([]string, 0)
	for _, c := range line {
		asString := string(c)
		if isOpening(c) {
			stack = append(stack, asString)
		} else {
			_, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
	}
	return stack
}

func repairIncomplete(line string) int {
	stack := getStack(line)	
	score := 0
	for i:=len(stack)-1; i>=0;i-- {
		char := stack[i]
		score *= 5
		score += getScore(char)
	}
	return score
}

func main() {
	dat := readData("input.txt")
	incompletes := filterIncomplete(dat)
	scores := make([]int, 0)
	for _, inc := range incompletes {
		scores = append(scores, repairIncomplete(inc))
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores) / 2])

	fmt.Println("finished")
}
