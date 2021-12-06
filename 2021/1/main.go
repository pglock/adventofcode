package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		num, _ := strconv.Atoi(lineStr)
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {
	dat := readData("input.txt")
	i := 1
	cnt := 0
	for i < len(dat) {
		if dat[i] > dat[i-1] {
			cnt += 1
		}
		i += 1
	}
	fmt.Println(cnt)
}
