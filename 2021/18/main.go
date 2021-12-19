package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type snailNumber struct {
	val int
	depth int
}

func readData(filename string) ([][]snailNumber, []string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	numbers := make([][]snailNumber, 0)
	lines := make([]string, 0)
	for scanner.Scan() {
		lineStr := scanner.Text()
		lines = append(lines, lineStr)
		number := make([]snailNumber, 0)
		depth := 0
		for _, char := range lineStr {
			switch char {
			case '[':
				depth++
				continue
			case ']':
				depth--
				continue
			case ',':
				continue
			default:
				num, _ := strconv.Atoi(string(char))
				number = append(number, snailNumber{num, depth})
			}
		}
		numbers = append(numbers, number)
	}
	return numbers, lines
}

func reduce(number []snailNumber) []snailNumber {
	changed := true
	for changed {
		changed = false
		for index:= range number {
			if number[index].depth >= 5 {
				if index > 0 {
					number[index-1].val += number[index].val
				}
				if index + 2 < len(number) {
					number[index +2].val += number[index+1].val
				}
				number[index].val = 0
				number[index].depth--
				number = append(number[:index+1], number[index+2:]...)
				changed = true
				break
			}
		}

		if changed {
			continue
		}

		for index := range number {
			if number[index].val > 9 {
				left := number[index].val / 2
				right := left
				if number[index].val % 2 != 0 {
					right += 1
				}
				newDepth := number[index].depth + 1
				sleft := snailNumber{left, newDepth}
				sright := snailNumber{right, newDepth}
				number = append(number[:index], append([]snailNumber{sleft, sright}, number[index+1:]...)...)
				changed = true
				break
			}
		}
	}

	return number
}

func add(left []snailNumber, right []snailNumber) []snailNumber {
	newNumber := append(left, right...)

	for i := range newNumber {
		newNumber[i].depth++
	}
	return newNumber
}

func magnitude(number []snailNumber) int {
	for depth:=4; depth>0; depth-=1 {
		newNumber := make([]snailNumber, 0)
		for i:=0;i<len(number);i++ {
			if number[i].depth == depth {
				val := 3 * number[i].val + 2 * number[i+1].val
				newNumber = append(newNumber, snailNumber{val, depth-1})
				i += 1
			} else {
				newNumber = append(newNumber, number[i])
			}
		}
		number = newNumber
	}
	return number[0].val
}

func main() {
	numbers, _ := readData("input.txt")
	number := numbers[0]
	for i:=1;i<len(numbers);i++ {
		number = add(number, numbers[i])
		number = reduce(number)
	}
	fmt.Println(magnitude(number))
}
