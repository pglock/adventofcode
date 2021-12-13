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

type fold struct {
	axis string
	value int
}

type point struct {
	row int
	col int
}

func readData(filename string) ([]point, []fold) {
	file, err := os.Open(filename)
	check(err)
	folds := make([]fold, 0)
	points := make([]point, 0)
	scanner := bufio.NewScanner(file)
	atFolds := false
	for scanner.Scan() {
		lineStr := scanner.Text()
		if len(strings.Fields(lineStr)) == 0 {
			atFolds = true
			continue
		}
		if atFolds {
			// read folding instructions
			words := strings.Fields(lineStr)
			splitted := strings.Split(words[2], "=")
			axis := splitted[0]
			value, _ := strconv.Atoi(splitted[1])
			folds = append(folds, fold{axis, value})
		} else {
			// read array data
			splitted := strings.Split(lineStr, ",")
			col, _ := strconv.Atoi(splitted[0])
			row, _ := strconv.Atoi(splitted[1])
			points = append(points, point{row, col})
		}


	}
	return points, folds
}

func contains(points []point, p point) bool {
	for _, pp := range points {
		if pp.row == p.row && pp.col == p.col {
			return true
		}
	}
	return false
}

func unique(points []point) []point {
	newPoints := make([]point, 0)
	for _, p := range points {
		if !contains(newPoints, p) {
			newPoints = append(newPoints, p)
		}
	}
	return newPoints
}

func horizontalFold(points []point, currentFold fold) []point {
	newPoints := make([]point, 0)
	for _, p := range points {
		if p.row < currentFold.value {
			newPoints = append(newPoints, p)
		} else if p.row > currentFold.value {
			newRow := 2 * currentFold.value - p.row
			newP := point{row: newRow, col: p.col}
			newPoints = append(newPoints, newP)
		}
	}
	newPoints = unique(newPoints)
	return newPoints
}

func verticalFold(points []point, currentFold fold) []point {
	newPoints := make([]point, 0)
	for _, p := range points {
		if p.col < currentFold.value {
			newPoints = append(newPoints, p)
		} else if p.col > currentFold.value {
			newCol := 2 * currentFold.value - p.col
			newP := point{row: p.row, col: newCol}
			newPoints = append(newPoints, newP)
		}
	}
	newPoints = unique(newPoints)
	return newPoints
}

func folding(points []point, folds []fold) []point {
	for _, currentFold := range folds {
		if currentFold.axis == "y" {
			points = horizontalFold(points, currentFold)
		} else {
			points = verticalFold(points, currentFold)
		}
	}
	return points
}

func prettyPrint(points []point) {
	// init array
	maxRow := -1
	maxCol := -1
	for _, p := range points {
		if p.row > maxRow {
			maxRow = p.row
		}
		if p.col > maxCol {
			maxCol = p.col
		}
	}
	array := make([][]string, 0)
	for i:=0; i<=maxRow; i++{
		currentCol := make([]string, maxCol+1)
		for j:=0;j<=maxCol;j++{
			currentCol[j] = "."
		}
		array = append(array, currentCol)
	}
	// write points
	for _, p := range points {
		array[p.row][p.col] = "#"
	}

	for _, row := range array {
		fmt.Println(row)
	}
}

func main() {
	points, folds := readData("input.txt")
	points = folding(points, folds)
	fmt.Println("n points: ", len(points))
	prettyPrint(points)
	fmt.Println("finished")
}
