package main

import (
	// "fmt"
	// "strconv"
	"strings"
)

type SlopeMap []string

func (m SlopeMap) CountTrees(dx, dy int) (count int) {
	var x, y int

	mapWidth := len(m[0])
	mapHeight := len(m)

	for y = 1; y <= mapHeight-1; y += dy {
		x = (x + dx) % mapWidth

		if m[y][x] == '#' {
			count++
		}
	}

	return
}

func ParseInput(input string) (res SlopeMap) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, line)
	}

	return
}
