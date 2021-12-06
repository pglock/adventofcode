package main

import (
	"bufio"
	"fmt"
	"math"
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
	data := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		row := make([]int, 0)
		for i := 0; i < len(lineStr); i++ {
			num, _ := strconv.Atoi(string(lineStr[i]))
			row = append(row, num)
		}
		data = append(data, row)
	}
	return data
}

func most_common(data [][]int, tie int) []int {
	most := make([]int, len(data[0]))

	half := float64(len(data)) / 2
	for j := 0; j < len(data[0]); j++ {
		sum := 0.0
		for i := 0; i < len(data); i++ {
			sum += float64(data[i][j])
		}
		if sum < half {
			most[j] = 0
		}
		if sum > half {
			most[j] = 1
		}
		if sum == half {
			most[j] = tie
		}

	}
	return most
}

func decimal(binary []int) int {
	length := len(binary)
	sum := 0.0
	for idx, v := range binary {
		if v == 1 {
			sum += math.Pow(2, float64(length-(idx+1)))
		}
	}
	return int(sum)
}

func oxygen(data [][]int, column int) int {
	most := most_common(data, 1)

	// filter out rows until one is left
	// iterator over most and check according column in each row, that's not yet filtered out
	// if they don't match filter out this row
	// return decimal
	value := most[column]
	newRows := make([][]int, 0)
	for i, _ := range data {
		if data[i][column] == value {
			newRows = append(newRows, data[i])
		}
	}
	if len(newRows) == 1 {
		return decimal(newRows[0])
	} else if column < len(newRows[0])-1 {
		return oxygen(newRows, column+1)
	}
	return -1
}

func co2(data [][]int, column int) int {
	most := most_common(data, 1)

	value := most[column]
	newRows := make([][]int, 0)
	for i, _ := range data {
		if data[i][column] != value {
			newRows = append(newRows, data[i])
		}
	}
	if len(newRows) == 1 {
		return decimal(newRows[0])
	} else if column < len(newRows[0])-1 {
		return co2(newRows, column+1)
	}
	return -1
}

func main() {
	dat := readData("input.txt")

	fmt.Println(oxygen(dat, 0))
	fmt.Println(co2(dat, 0))
}
