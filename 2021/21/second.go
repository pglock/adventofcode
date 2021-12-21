package main

import (
	"fmt"
)

type player struct {
	pos int
	score int
}

type universe struct {
	p1 player
	p2 player
	p1Turn bool
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

func step(universes map[universe]int) (map[universe]int, int, int) {
	newUniverses := make(map[universe]int)
	p1wins := 0
	p2wins := 0
	for uni, cnt := range universes {
		newUnis, p1w, p2w := turn(uni)
		p1wins += cnt * p1w
		p2wins += cnt * p2w
		for newU, newCnt := range newUnis {
			newUniverses[newU] += newCnt * cnt
		}
	}
	return newUniverses, p1wins, p2wins
}

func turn(uni universe) (map[universe]int, int, int) {
	p1Win := 0
	p2Win := 0
	newUnis := make(map[universe]int)
	possibleRolls := map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
	for val, cnt := range possibleRolls {
		tmpUni := uni
		if tmpUni.p1Turn {
			tmpUni.p1.pos = calcPos(tmpUni.p1.pos, val)
			tmpUni.p1.score += tmpUni.p1.pos
			if tmpUni.p1.score >= 21 {
				p1Win += cnt
				continue
			}
		} else {
			tmpUni.p2.pos = calcPos(tmpUni.p2.pos, val)
			tmpUni.p2.score += tmpUni.p2.pos
			if tmpUni.p2.score >= 21 {
				p2Win += cnt
				continue
			}
		}
		tmpUni.p1Turn = !tmpUni.p1Turn
		newUnis[tmpUni] += cnt
	}
	return newUnis, p1Win, p2Win
}

func calcPos(pos int, add int) int {
	return (pos + add - 1) % 10 + 1
}

func diracDice(player1 player, player2 player) int {
	p1total := 0
	p2total := 0
	universes := make(map[universe]int)
	start := universe{player1, player2, true}
	universes[start] = 1
	p1Wins := 0
	p2Wins := 0
	for len(universes) > 0 {
		universes, p1Wins, p2Wins = step(universes)
		p1total += p1Wins
		p2total += p2Wins
	}
	if p1total > p2total {
		return p1total
	}
	return p2total
}

func main() {
	p1 := player{7, 0}
	p2 := player{9, 0}
	fmt.Println(diracDice(p1, p2))
}
