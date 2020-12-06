package main

import (
	// "fmt"
	// "sort"
	// "strconv"
	"strings"
)

func ParseInput(input string) (res []CustomGroup) {
	var currentGroup CustomGroup
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			res = append(res, currentGroup)
			currentGroup = CustomGroup{}
		} else {
			currentGroup.RawInputs = append(currentGroup.RawInputs, line)
		}
	}

	return append(res, currentGroup)
}

func SumAllQuestions(groups []CustomGroup, countFn func(CustomGroup) int) (count int) {
	for _, group := range groups {
		count += countFn(group)
	}

	return
}

type CustomGroup struct {
	RawInputs []string
}

func (g CustomGroup) CountAnyoneYeses() int {
	set := make(map[rune]bool)
	for _, input := range g.RawInputs {
		for _, ch := range input {
			set[ch] = true
		}
	}

	return len(set)
}

// func (p BoardingPass) GetRow() int {
// 	rowSpec := p.AirlineSpec[0:7]
// 	return p.convertBinary(rowSpec, 'F', 'B')
// }

// func (p BoardingPass) GetCol() int {
// 	columnSpec := p.AirlineSpec[7:10]
// 	return p.convertBinary(columnSpec, 'L', 'R')
// }

// func (p BoardingPass) GetSeatId() int {
// 	return p.GetRow()*8 + p.GetCol()
// }

// func (p BoardingPass) convertBinary(spec string, lowCh, highCh byte) int {
// 	var binaryStr strings.Builder
// 	for _, ch := range spec {
// 		switch byte(ch) {
// 		case lowCh:
// 			binaryStr.WriteString("0")
// 		case highCh:
// 			binaryStr.WriteString("1")
// 		}
// 	}

// 	num, _ := strconv.ParseInt(binaryStr.String(), 2, 64)

// 	return int(num)
// }

// func FindMaxSeatId(passes []BoardingPass) (res int) {
// 	for _, p := range passes {
// 		if seatId := p.GetSeatId(); seatId > res {
// 			res = seatId
// 		}
// 	}

// 	return
// }

// func AllSeatIds(passes []BoardingPass) (res []int) {
// 	for _, p := range passes {
// 		res = append(res, p.GetSeatId())
// 	}
// 	sort.Ints(res)

// 	return
// }

// func FindMissingNum(sortedSeatIds []int) int {
// 	firstId := sortedSeatIds[0]
// 	for i, p := range sortedSeatIds {
// 		if p-i > firstId {
// 			return p - 1
// 		}
// 	}

// 	return 0
// }
