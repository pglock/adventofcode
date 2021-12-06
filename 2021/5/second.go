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

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type field struct {
	data [][]int
}

func readData(filename string) []line {
	file, err := os.Open(filename)
	check(err)
	lines := make([]line, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		w := strings.Fields(lineStr)
		start := strings.Split(w[0], ",")
		end := strings.Split(w[2], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		lines = append(lines, line{x1, y1, x2, y2})

	}
	return lines
}

func getMax(lines []line) (int, int) {
	maxX := 0
	maxY := 0
	for _, line := range lines {
		if maxX < line.x1 {
			maxX = line.x1
		}
		if maxY < line.y1 {
			maxY = line.y1
		}
		if maxX < line.x2 {
			maxX = line.x2
		}
		if maxY < line.y2 {
			maxY = line.y2
		}
	}
	return maxX + 1, maxY + 1
}

func initField(maxX int, maxY int) field {
	data := make([][]int, 0)
	for i := 0; i < maxY; i++ {
		data = append(data, make([]int, maxX))
	}
	return field{data}
}

func getStartAndStop(line line) (int, int, int, int) {
	sY := line.y1
	eY := line.y2
	if line.y1 > line.y2 {
		sY = line.y2
		eY = line.y1
	}
	sX := line.x1
	eX := line.x2
	if line.x1 > line.x2 {
		sX = line.x2
		eX = line.x1
	}
	return sX, sY, eX, eY
}

func addDiagLine(field *field, line line) {
	sX, sY, eX, _ := getStartAndStop(line)
	diff := eX - sX
	sigmaX := 1
	if sX != line.x1 {
		sigmaX = -1
	}
	sigmaY := 1
	if sY != line.y1 {
		sigmaY = -1
	}
	for i := 0; i <= diff; i++ {
		field.data[line.y1+sigmaY*i][line.x1+sigmaX*i] += 1
	}
}

func addNormalLine(field *field, line line) {
	sX, sY, eX, eY := getStartAndStop(line)
	for i := sY; i <= eY; i++ {
		for j := sX; j <= eX; j++ {
			field.data[i][j] += 1
		}
	}
}

func addLine(field *field, line line) {
	sX, sY, eX, eY := getStartAndStop(line)
	if line.x1 == line.x2 || line.y1 == line.y2 {
		addNormalLine(field, line)
	} else if eY-sY == eX-sX {
		addDiagLine(field, line)
	}
}

func generateField(lines []line) field {
	maxX, maxY := getMax(lines)
	field := initField(maxX, maxY)
	for _, line := range lines {
		addLine(&field, line)
	}
	return field
}

func sumField(field field) int {
	sum := 0
	for _, row := range field.data {
		for _, val := range row {
			if val > 1 {
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	lines := readData("input.txt")
	field := generateField(lines)
	result := sumField(field)
	fmt.Println(result)
	fmt.Println("finished")
}
