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
	cnt := 0
	for i := 1; i < len(dat)-2; i++ {
		sum1 := dat[i-1] + dat[i] + dat[i+1]
		sum2 := dat[i] + dat[i+1] + dat[i+2]
		if sum2 > sum1 {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
