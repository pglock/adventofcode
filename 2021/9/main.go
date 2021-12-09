package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func readData(filename string) [][]int {
	file, err := os.Open(filename)
	check(err)
	array := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		row := make([]int, 0)
		for _, c := range lineStr {
			num, _ := strconv.Atoi(string(c))
			row = append(row, num)
		}
		array = append(array, row)
	}
	return array
}

func getNeighbors(array [][]int, row int, col int) []int {
	neighbors := make([]int,0)
	if row > 0 {
		neighbors = append(neighbors, array[row-1][col])
	}
	if col > 0 {
		neighbors = append(neighbors, array[row][col -1])
	}
	if row < len(array) - 1 {
		neighbors = append(neighbors, array[row +1][col])
	}
	if col < len(array[0]) - 1 {
		neighbors = append(neighbors, array[row][col +1])
	}
	return neighbors
}

func isMinimum(list []int, val int) bool {
	result := true
	for _, el := range list {
		if el <= val {
			result = false
		}
	}
	return result
}


func findMinimums(array [][]int) []int {
	minimums := make([]int, 0)
	for i, row := range array {
		for j, val := range row {
			neighbors := getNeighbors(array, i, j)
			if isMinimum(neighbors, val) {
				minimums = append(minimums, val)
			}
		}
	}
	return minimums
}

func sumRisk(array []int) int {
	sum := 0
	for _, v := range array {
		sum += v + 1
	}
	return sum
}

func main() {
	dat := readData("input.txt")
	for _, row := range dat  {
		fmt.Println(row)
	}
	fmt.Println(sumRisk(findMinimums(dat)))
	fmt.Println("finished")
}
