package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	value     int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(filename string) []command {
	file, err := os.Open(filename)
	check(err)
	commands := []command{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		dir := strings.Split(lineStr, " ")[0]
		val := strings.Split(lineStr, " ")[1]
		num, _ := strconv.Atoi(val)
		commands = append(commands, command{direction: dir, value: num})
	}
	return commands
}

func main() {
	dat := readData("input.txt")
	x := 0
	y := 0
	for i := 0; i < len(dat); i++ {
		if dat[i].direction == "forward" {
			x += dat[i].value
		}
		if dat[i].direction == "up" {
			y -= dat[i].value
		}
		if dat[i].direction == "down" {
			y += dat[i].value
		}
	}
	fmt.Println(x)
	fmt.Println(y)
}
