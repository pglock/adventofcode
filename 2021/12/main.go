package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

type node struct {
	id string
	friends map[string]*node
}

func readData(filename string) map[string][]string {
	file, err := os.Open(filename)
	check(err)
	allNodes := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		splitted := strings.Split(lineStr, "-")
		if _, ok := allNodes[splitted[0]]; !ok {
			allNodes[splitted[0]] = make([]string, 0)
		}
		allNodes[splitted[0]] = append(allNodes[splitted[0]], splitted[1])
		if _, ok := allNodes[splitted[1]]; !ok {
			allNodes[splitted[1]] = make([]string, 0)
		}
		allNodes[splitted[1]] = append(allNodes[splitted[1]], splitted[0])
	}
	return allNodes
}

func contains( path []string, node string) bool {
	for _, currentNode := range path {
		if currentNode == node {
			return true
		}
	}
	return false
}

func bfs(graph map[string][]string, start string) [][]string {
	paths := make([][]string, 0)
	queue := make([][]string, 0)
	queue = append(queue, []string{start})

	for ok:=true; ok; ok = len(queue) > 0 {
		currentPath := queue[0]
		queue = queue[1:]
		lastNode := currentPath[len(currentPath) - 1]
		if lastNode == "end" {
			paths = append(paths, currentPath)
			continue
		}

		friends, hasFriends := graph[lastNode]
		if hasFriends {
			for _, friend := range friends {
				newPath := make([]string, len(currentPath))
				copy(newPath, currentPath)
				if friend == strings.ToUpper(friend) || !contains(currentPath, friend) {
					newPath = append(newPath, friend)
					queue = append(queue, newPath)
				}
			}
		}

	}

	return paths
}

func main() {
	dat := readData("input.txt")
	paths := bfs(dat, "start")
	fmt.Println(len(paths))
	fmt.Println("finished")
}
