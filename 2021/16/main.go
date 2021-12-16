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

func parsePackage(data string, head int) (int, int) {
	oldHead := head
	version := parseInt(data[head:head+3])
	head += 3
	typeId := parseInt(data[head:head+3])
	head += 3
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
	} else {
		lengthType := parseInt(data[head:head+1])
		head += 1
		if lengthType == 0 {
			totalLength := parseInt(data[head:head+15])
			head += 15
			for totalLength > 0 {
				subVersion, subHead := parsePackage(data, head)
				head += subHead
				totalLength -= subHead
				version += subVersion
			}
		} else {
			nSubPackages := parseInt(data[head:head+11])
			head += 11
			for i:=0;i<nSubPackages; i++{
				subVersion, subHead := parsePackage(data, head)
				head += subHead
				version += subVersion
			}
		}

	}
	return version, head-oldHead
}

func parseInt(data string) int {
	version, _ := strconv.ParseInt(data, 2, 64)
	return int(version)
}

func main() {
	binary := readData("input.txt")
	version, head := parsePackage(binary, 0)
	fmt.Println("version, head", version, head)
}
