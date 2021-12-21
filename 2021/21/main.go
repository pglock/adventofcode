package main

import (
	"fmt"
)

type player struct {
	pos int
	score int
}

func deterministicDice(player1 player, player2 player) int {
	currentPlayer := &player1
	dice := 1
	for ok := true; ok; ok = player1.score < 1000 && player2.score < 1000 {
		realDice := (dice - 1) % 100 + 1
		currentPlayer.pos = (currentPlayer.pos + realDice + realDice + 1 + realDice + 2 - 1) % 10 + 1
		currentPlayer.score += currentPlayer.pos
		if currentPlayer == &player1 {
			currentPlayer = &player2
		} else {
			currentPlayer =&player1
		}
		dice += 3
	}
	fmt.Println("score", currentPlayer.score)
	fmt.Println("dice", dice)
	return currentPlayer.score * (dice -1)
}

func main() {
	p1 := player{7, 0}
	p2 := player{9, 0}
	fmt.Println(deterministicDice(p1, p2))
}
