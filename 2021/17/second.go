package main

import (
	"fmt"
)

type Target struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

type velocity struct {
	x int
	y int
}

func contains(list []velocity, vel velocity) bool {
	for _, el := range list {
		if el == vel {
			return true
		}
	}
	return false
}

func reaches(target Target, vel velocity) bool {
	startPos := velocity{0, 0}
	for ok :=true; ok; ok = !overshot(target, startPos, vel) {
		startPos.x += vel.x
		startPos.y += vel.y
		vel.y -= 1
		if vel.x > 0 {
			vel.x -= 1
		} else if vel.x < 0 {
			vel.x += 1
		}
		if check(target, startPos) {
			return true
		}
	}
	return false
}

func overshot(target Target, pos velocity, vel velocity) bool {
	if pos.x > target.xMax || pos.y < target.yMin {
		return true
	}
	if (pos.x < target.xMin || pos.x > target.xMax) && vel.x == 0 {
		return true
	}
	return false
}

func check(target Target, pos velocity) bool {

	if pos.x >= target.xMin && pos.x <= target.xMax {
		if pos.y >= target.yMin && pos.y <= target.yMax {
			return true
		}
	}
	return false
}

func main() {
	target := Target{111, 161, -154, -101}
	allMatches := make([]velocity, 0)

	for currentX:=1; currentX < 1000; currentX++ {
		for currentY:=-160; currentY < 160; currentY++ {
			vel := velocity{currentX, currentY}
			if reaches(target, vel) && !contains(allMatches, vel) {
				allMatches = append(allMatches, vel)
			}
		}
	}

	fmt.Println(len(allMatches))
}
