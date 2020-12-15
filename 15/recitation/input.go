package recitation

import (
	"fmt"
	// "sort"
	"strconv"
	"strings"
)

var _ = fmt.Println

type MemoryGame struct {
	Numbers    []int
	LastSpoken map[int]int
}

func (g *MemoryGame) TakeTurnsUntil(iteration int) {
	for len(g.Numbers) < iteration {
		g.NextTurn()
	}
}

func (g *MemoryGame) NextTurn() {
	prevTurn := len(g.Numbers) - 1
	prevNum := g.Numbers[prevTurn]
	if lastSpoken, exists := g.LastSpoken[prevNum]; exists && lastSpoken < prevTurn {
		g.Numbers = append(g.Numbers, prevTurn-lastSpoken)
	} else {
		g.Numbers = append(g.Numbers, 0)
	}
	g.LastSpoken[prevNum] = prevTurn
}

func (g *MemoryGame) LastNumber() int {
	return g.Numbers[len(g.Numbers)-1]
}

func ParseInput(input string) *MemoryGame {
	res := MemoryGame{
		Numbers:    []int{},
		LastSpoken: make(map[int]int),
	}

	for idx, numStr := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(numStr)
		res.Numbers = append(res.Numbers, num)
		res.LastSpoken[num] = idx
	}

	return &res
}
