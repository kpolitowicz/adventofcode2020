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
