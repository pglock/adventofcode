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

func readData(filename string) [9]int {
	file, err := os.Open(filename)
	check(err)
	var numbers [9]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		splits := strings.Split(lineStr, ",")
		for _, split := range splits {
			num, _ := strconv.Atoi(split)
			numbers[num] += 1
		}

	}
	return numbers
}

func days(fishs [9]int, nDays int) [9]int {
	for day := 0; day < nDays; day++ {
		fishs[(day+7)%9] += fishs[day%9]
	}
	return fishs
}

func sum(fishs [9]int) int {
	sum := 0
	for _, v := range fishs {
		sum += v
	}
	return sum
}

func printFishs(fishs [9]int) {
	for i, f := range fishs {
		fmt.Println(i, f)
	}
}

func main() {
	dat := readData("input.txt")
	printFishs(dat)
	resState := days(dat, 256)
	fmt.Println(sum(resState))
	fmt.Println("finished")
}
