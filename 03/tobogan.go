package main

import (
	// "fmt"
	// "strconv"
	"strings"
)

type SlopeMap []string

func (m SlopeMap) CountTrees() (count int) {
	var x, y int
	mapWidth := len(m[0])

	for y = 1; y <= len(m)-1; y++ {
		x = (x + 3) % mapWidth

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
