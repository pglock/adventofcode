package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	numbers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		splits := strings.Split(lineStr, ",")
		for _, split := range splits {
			num, _ := strconv.Atoi(split)
			numbers = append(numbers, num)
		}

	}
	return numbers
}

func day(fishs []int) []int {
	toAdd := 0
	for i, _ := range fishs {
		fishs[i] -= 1
		if fishs[i] == -1 {
			toAdd += 1
			fishs[i] = 6
		}
	}
	for i := 0; i < toAdd; i++ {
		fishs = append(fishs, 8)
	}
	return fishs
}

func main() {
	dat := readData("input.txt")
	for i := 0; i < 80; i++ {
		dat = day(dat)
	}
	fmt.Println(len(dat))
	fmt.Println("finished")
}
