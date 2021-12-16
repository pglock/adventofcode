package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
)

var hex2bin = map[string]string{
"0" : "0000",
"1" : "0001",
"2" : "0010",
"3" : "0011",
"4" : "0100",
"5" : "0101",
"6" : "0110",
"7" : "0111",
"8" : "1000",
"9" : "1001",
"A" : "1010",
"B" : "1011",
"C" : "1100",
"D" : "1101",
"E" : "1110",
"F" : "1111",
}

func readData(filename string) string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	binary := ""
	for scanner.Scan() {
		lineStr := scanner.Text()
		for _, char := range lineStr {
			binary = binary + hex2bin[string(char)]
		}
	}
	return binary
}

func parsePackage(data string, head int) (int, int, int) {
	oldHead := head
	version := parseInt(data[head:head+3])
	head += 3
	typeId := parseInt(data[head:head+3])
	head += 3
	value := 0
	if typeId == 4 {
		// todo read actual data
		binary := ""
		last:=1
		for ok := true; ok; ok = last != 0 {
			last = parseInt(data[head: head+1])
			head += 1
			binary += data[head:head+4]
			head += 4
		}
		value = parseInt(binary)
	} else {
		lengthType := parseInt(data[head:head+1])
		head += 1
		subValues := make([]int, 0)
		if lengthType == 0 {
			totalLength := parseInt(data[head:head+15])
			head += 15
			for totalLength > 0 {
				subVersion, subHead, subValue := parsePackage(data, head)
				head += subHead
				totalLength -= subHead
				version += subVersion
				subValues = append(subValues, subValue)
			}
		} else {
			nSubPackages := parseInt(data[head:head+11])
			head += 11
			for i:=0;i<nSubPackages; i++{
				subVersion, subHead, subValue := parsePackage(data, head)
				head += subHead
				version += subVersion
				subValues = append(subValues, subValue)
			}
		}

		switch typeId {
		case 0 :
			value = sum(subValues)
		case 1 :
			value = prod(subValues)
		case 2 :
			value = min(subValues)
		case 3 :
			value = max(subValues)
		case 5:
			value = greaterThan(subValues)
		case 6:
			value = lessThan(subValues)
		case 7:
			value = equals(subValues)
		}

	}
	return version, head-oldHead, value
}

func sum(values []int) int {
	s := 0
	for _, v := range values {
		s += v
	}
	return s
}

func prod(values []int) int {
	s := 1
	for _, v := range values {
		s *= v
	}
	return s
}

func min(values []int) int {
	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func max(values []int) int {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func greaterThan(values []int) int {
	if values[0] > values[1] {
		return 1
	} 
	return 0
}

func lessThan(values []int) int {
	if values[0] < values[1] {
		return 1
	} 
	return 0
}

func equals(values []int) int {
	if values[0] == values[1] {
		return 1
	} 
	return 0
}

func parseInt(data string) int {
	version, _ := strconv.ParseInt(data, 2, 64)
	return int(version)
}

func main() {
	binary := readData("input.txt")
	version, head, value := parsePackage(binary, 0)
	fmt.Println("version, head, value", version, head, value)
}
