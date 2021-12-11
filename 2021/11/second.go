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

func readData(filename string) [][]int {
	file, err := os.Open(filename)
	check(err)
	array := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		row := make([]int, 0)
		for _, char := range lineStr {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		array = append(array, row)
	}
	return array
}

func isValid(row int, col int, nRows int, nCols int) bool {
	if row >= 0 && row < nRows {
		if col >= 0 && col < nCols {
			return true
		}
	}
	return false
}

func getNeighbors(array [][]int, row int, col int) [][]int {
	neighbors := make([][]int, 0)
	nRows := len(array)
	nCols := len(array[0])
	for i:=-1; i<=1; i++ {
		for j:=-1;j<=1;j++ {
			if i==0 && j==0 {
				continue
			}
			if isValid(row + i, col + j, nRows, nCols){
				if array[row+i][col+j] >= 0 {
					neighbors = append(neighbors, []int{row + i, col + j})
				}
			}
		}
	}
	return neighbors
}

func step(array [][]int) int {
	cnt := 0
	to_visit := make([][]int, 0)
	for i, row := range array {
		for j := range row {
			array[i][j] += 1
			if array[i][j] > 9 {
				array[i][j] = -1
				cnt += 1
				to_visit = append(to_visit, getNeighbors(array, i, j)...)
			}
		}
	}
	for ok := len(to_visit) > 0; ok; ok = len(to_visit) > 0 {
		point := to_visit[0]
		to_visit = to_visit[1:]
		if array[point[0]][point[1]] > 0 {
			array[point[0]][point[1]] += 1
			if array[point[0]][point[1]] > 9 {
				array[point[0]][point[1]] = -1
				cnt += 1
				to_visit = append(to_visit, getNeighbors(array, point[0], point[1])...)
			}
		}
	}
	for i, row := range array {
		for j, val := range row {
			if val == -1 {
				array[i][j] = 0
			}
		}
	}
	return cnt
}

func main() {
	dat := readData("input.txt")
	total := 0
	nStep := 0
	nOcto := len(dat) * len(dat[0])
	for i:=true; i; i = total != nOcto {
		cnt := step(dat)
		total = cnt
		nStep += 1
	}
	fmt.Println(nStep)
	fmt.Println("finished")
}
