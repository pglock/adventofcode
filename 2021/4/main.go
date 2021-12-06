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

type board struct {
	fields [5][5]int
}

func readCommands(line string) []int {
	numbers := make([]int, 0)
	for _, val := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}
	return numbers
}

func readData(filename string) ([]int, []board) {
	file, err := os.Open(filename)
	check(err)
	boards := make([]board, 0)
	var numbers []int
	scanner := bufio.NewScanner(file)
	currentBoard := board{}
	currentRow := 0
	currentLine := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		words := strings.Fields(lineStr)
		if currentLine == 0 {
			// first line is list of drawn numbers
			numbers = readCommands(lineStr)
		} else if len(words) == 0 {
			// skip empty rows, if data is in the current board save it
			if currentRow > 0 {
				boards = append(boards, currentBoard)
				currentBoard = board{}
				currentRow = 0
			} else {
				currentLine += 1
				continue
			}
		} else {
			// write values to boards
			for i, val := range words {
				num, _ := strconv.Atoi(val)
				currentBoard.fields[currentRow][i] = num
			}
			currentRow += 1
		}
		currentLine += 1
	}
	// save last board
	boards = append(boards, currentBoard)
	return numbers, boards
}

func mark(board *board, value int) {
	for i, row := range board.fields {
		for j, val := range row {
			if val == value {
				board.fields[i][j] = -1
			}
		}
	}
}

func checkRow(row [5]int) bool {
	for _, val := range row {
		if val != -1 {
			return false
		}
	}
	return true
}

func checkCol(fields [5][5]int, col int) bool {
	for _, row := range fields {
		if row[col] != -1 {
			return false
		}
	}
	return true
}

func checkFinished(board board) bool {
	for _, row := range board.fields {
		if checkRow(row) {
			return true
		}
	}
	for i := 0; i < 5; i++ {
		if checkCol(board.fields, i) {
			return true
		}
	}
	return false
}

func sumBoard(board board) int {
	sum := 0
	for _, row := range board.fields {
		for _, val := range row {
			if val != -1 {
				sum += val
			}
		}
	}
	return sum
}

func main() {
	commands, boards := readData("input.txt")
	fmt.Println(commands)
	won := false
	for _, command := range commands {
		for i, _ := range boards {
			mark(&boards[i], command)
			if checkFinished(boards[i]) {
				boardSum := sumBoard(boards[i])
				fmt.Println("Result", boardSum*command)
				won = true
				break
			}
		}
		if won {
			break
		}
	}
	fmt.Println("finished")

}
