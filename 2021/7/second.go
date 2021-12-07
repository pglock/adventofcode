package main

import (
	"bufio"
	"fmt"
	"math"
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

func getMax(numbers []int) int {
	max := -1
	for _, v := range numbers {
		if max < v {
			max = v
		}
	}
	return max
}

func calcDistance(numbers []int, target int) int {
	sum := 0
	for _, val := range numbers {
		steps := int(math.Abs(float64(target-val))) + 1
		sum += (steps*steps - steps) / 2
	}
	return sum
}

func minCosts(numbers []int) int {
	max := getMax(numbers)
	min := -1
	for i := 0; i < max; i++ {
		tmp := calcDistance(numbers, i)
		if min == -1 || min > tmp {
			min = tmp
		}
	}
	return min
}

func main() {
	dat := readData("input.txt")
	fmt.Println(minCosts(dat))
	fmt.Println("finished")
}
