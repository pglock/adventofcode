package main

import (
	"container/heap"
	"strconv"
	"os"
	"bufio"
	"fmt"
)

type position struct {
	row int
	col int
}

type Item struct {
	pos position
	riskLevel int

	index int // needed by heap
}

// PriorityQueue definition
type PriorityQueue []Item
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	
	return pq[i].riskLevel < pq[j].riskLevel
	
}
func (pq PriorityQueue) Swap(i, j int) {
	
	pq[i], pq[j] = pq[j], pq[i]
	
	pq[i].index = i
	
	pq[j].index = j
	
}

func (pq *PriorityQueue) Push(x interface{}) {
	
	n := len(*pq)
	
	item := x.(Item)
	
	item.index = n
	
	*pq = append(*pq, item)
	
}

func (pq *PriorityQueue) Pop() interface{} {
	
	old := *pq
	
	n := len(old)
	
	item := old[n-1]
	
	// old[n-1] = nil  // avoid memory leak
	
	item.index = -1 // for safety
	
	*pq = old[0 : n-1]
	
	return item
	
}

func readData(filename string) (map[position]int, int, int) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	array := make(map[position]int)
	i := 0
	nRows := 0
	nCols := 0
	for scanner.Scan() {
		lineStr := scanner.Text()
		for j, char := range lineStr {
			num, _ := strconv.Atoi(string(char))
			array[position{i, j}] = num
			nCols = j
		}
		i ++
		nRows = i
	}
	return array, nRows, nCols + 1
}

func neighbors(p position, nRows int, nCols int) []position {
	points := make([]position, 0)
	if p.row > 1 {
		points = append(points, position{row: p.row-1,col: p.col})
	}
	if p.row < nRows -1 {
		points = append(points, position{row: p.row+1,col: p.col})
	}
	if p.col > 1 {
		points = append(points, position{row: p.row,  col: p.col-1})
	}
	if p.col < nCols -1 {
		points = append(points, position{row: p.row,  col: p.col+1})
	}
	return points
}

func dijkstra(grid map[position]int, nRows int, nCols int) int {
	start := position{0,0}
	target := position{nRows - 1, nCols - 1}
	costSoFar := make(map[position]int)
	costSoFar[start] = 0
	cameFrom := make(map[position]position)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	pq.Push(Item{pos: start, riskLevel: 0})

	for pq.Len() > 0 {
		head := heap.Pop(&pq).(Item)
		if head.pos == target {
			break
		}

		for _, neighbor := range neighbors(head.pos, nRows, nCols) {
			newCost := costSoFar[head.pos] + grid[neighbor]
			if old, ok := costSoFar[neighbor]; !ok || old > newCost {
				costSoFar[neighbor] = newCost
				pq.Push(Item{pos: neighbor, riskLevel: newCost})
				cameFrom[neighbor] = head.pos
			}
		}
	}
	return costSoFar[target]
}

func buildPath(cameFrom map[position]position, target position) []position {
	path := make([]position, 0)
	current := target
	start := position{0, 0}
	for ok:=true;ok;ok = current != start {
		path = append([]position{current}, path ...)
		current = cameFrom[current]
	}
	return path
}

func main() {
	array, nRows, nCols := readData("input.txt")
	fmt.Println(dijkstra(array, nRows, nCols))
	fmt.Println(nRows, nCols)
}
