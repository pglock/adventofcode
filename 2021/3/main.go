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

func calcGamma(data [][]int) float64 {
	gamma := 0.0
	for j := 0; j < len(data[0]); j++ {
		sum := 0
		for i := 0; i < len(data); i++ {
			sum += data[i][j]
		}
		if sum > len(data)/2 {
			gamma += math.Pow(float64(2), float64(len(data[0])-1.0-j))
		}
	}
	return gamma
}

func calcEpsilon(data [][]int) float64 {
	eps := 0.0
	for j := 0; j < len(data[0]); j++ {
		sum := 0
		for i := 0; i < len(data); i++ {
			sum += data[i][j]
		}
		if sum <= len(data)/2 {
			eps += math.Pow(float64(2), float64(len(data[0])-1.0-j))
		}
	}
	return eps
}

func main() {
	dat := readData("input.txt")
	var gamma float64 = calcGamma(dat)
	var eps float64 = calcEpsilon(dat)
	fmt.Println(gamma)
	fmt.Println(eps)
	//fmt.Println(y)
}
