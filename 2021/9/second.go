package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

type point struct {
	row int
	col int
	val int
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

func getNeighbors(array [][]int, p point) []point {
	neighbors := make([]point, 0)
	if p.row > 0 {
		neighbors = append(neighbors, point{p.row-1, p.col, array[p.row-1][p.col]})
	}
	if p.col > 0 {
		neighbors = append(neighbors, point{p.row, p.col-1, array[p.row][p.col-1]})
	}
	if p.row < len(array)-1 {
		neighbors = append(neighbors, point{p.row+1, p.col, array[p.row+1][p.col]})
	}
	if p.col < len(array[0])-1 {
		neighbors = append(neighbors, point{p.row, p.col+1, array[p.row][p.col+1]})
	}
	return neighbors
}

func isMinimum(list []point, val int) bool {
	result := true
	for _, el := range list {
		if el.val <= val {
			result = false
		}
	}
	return result
}

func findMinimums(array [][]int) []point {
	minimums := make([]point, 0)
	for i, row := range array {
		for j, val := range row {
			p := point{i, j, val}
			neighbors := getNeighbors(array, p)
			if isMinimum(neighbors, val) {
				minimums = append(minimums, p)
			}
		}
	}
	return minimums
}

func inList(array []point, p point) bool {
	for _, pp := range array {
		if pp.row == p.row && pp.col == p.col {
			return true
		}
	}
	return false
}

func buildBasin(array [][]int, low point) []point {
	basin := make([]point, 0)
	to_visit := make([]point, 1)
	to_visit[0] = low
	p := point{}
	for ok := true; ok; ok = len(to_visit) > 0 {
		p, to_visit = to_visit[0], to_visit[1:]
		if !inList(basin, p) {
			basin = append(basin, p)
		}
		for _, neighbor := range getNeighbors(array, p) {
			if inList(basin, neighbor) || neighbor.val == 9 {
				continue
			}
			to_visit = append(to_visit, neighbor)
		}
	}
	return basin
}

func maxBasins(basins [][]point) int {
	lengths := make([]int, 0)
	for _, basin := range basins {
		lengths = append(lengths, len(basin))
	}
	sort.Ints(lengths)
	nBasins := len(basins)
	return lengths[nBasins - 1] * lengths[nBasins - 2] * lengths[nBasins - 3] 
}

func main() {
	dat := readData("input.txt")
	minimums := findMinimums(dat)
	basins := make([][]point, 0)
	for i:=0; i<len(minimums);i++ {
		basin := buildBasin(dat, minimums[i])
		basins = append(basins, basin)
	}
	fmt.Println(maxBasins(basins))
	fmt.Println("finished")
}
