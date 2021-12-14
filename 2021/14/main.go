package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

type kv struct {
	key string
	value int 
}

func readData(filename string) (map[string]int, map[string]string, string, string){
	file, err := os.Open(filename)
	check(err)
	rules := make(map[string]string)
	template := make(map[string]int)
	atRules := false
	scanner := bufio.NewScanner(file)
	var start string
	var end string
	for scanner.Scan() {
		lineStr := scanner.Text()
		if len(strings.Fields(lineStr)) == 0{
			atRules = true
			continue
		}
		if !atRules {
			rawTemplate := lineStr
			start = string(lineStr[0])
			end = string(lineStr[len(lineStr) - 1])
			// read template to pairs
			for i:=0; i < len(rawTemplate) - 1; i++ {
				pair := rawTemplate[i:i+2]
				if _, ok := template[pair]; !ok {
					template[pair] = 1
				} else {
					template[pair] += 1
				}
			}
		} else {
			words := strings.Fields(lineStr)
			rules[words[0]] = words[2]
		}
	}
	return template, rules, start, end 
}

func step(template map[string]int, rules map[string]string) map[string]int {
	newTemplate := make(map[string]int)
	for pair, count := range template {
		foundRule := false
		for rulePair, toAdd := range rules {
			if pair == rulePair {
				newTemplate[string(pair[0]) + toAdd] += count
				newTemplate[toAdd + string(pair[1])] += count
				foundRule = true
				
			} 
		}
		if !foundRule {
			newTemplate[pair] += count
		}
	}

	return newTemplate
}

func length(template map[string]int) int {
	total := 0
	for _, count := range template {
		total += count
	}
	return total + 1
}

func countChar(template map[string]int, start string, end string) int {
	var sorted []kv
	chars := make(map[string]int)
	for pair, count := range template {
		chars[string(pair[0])] += count
		chars[string(pair[1])] += count
	}
	for k, v := range chars {
		sorted = append(sorted, kv{k, v})
	}


	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].value < sorted[j].value
	})
	fmt.Println(sorted[len(sorted)-1], sorted[0])

	maxChar := sorted[len(sorted)-1]
	maxVal := maxChar.value / 2
	if maxChar.key == start || maxChar.key == end {
		maxVal += 1
	}
	minChar := sorted[0]
	minVal := minChar.value / 2
	if minChar.key == start || minChar.key == end {
		minVal += 1
	}
	return maxVal - minVal
}

func main() {
	template, rules, start, end := readData("input.txt")
	fmt.Println(length(template))
	for i:=0;i<10;i++ {
		template = step(template, rules)
	}
	fmt.Println(countChar(template, start, end))
	fmt.Println("finished")
}
