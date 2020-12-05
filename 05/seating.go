package main

import (
	// "fmt"
	"sort"
	"strconv"
	"strings"
)

type BoardingPass struct {
	AirlineSpec string
}

func (p BoardingPass) GetRow() int {
	rowSpec := p.AirlineSpec[0:7]
	return p.convertBinary(rowSpec, 'F', 'B')
}

func (p BoardingPass) GetCol() int {
	columnSpec := p.AirlineSpec[7:10]
	return p.convertBinary(columnSpec, 'L', 'R')
}

func (p BoardingPass) GetSeatId() int {
	return p.GetRow()*8 + p.GetCol()
}

func (p BoardingPass) convertBinary(spec string, lowCh, highCh byte) int {
	var binaryStr strings.Builder
	for _, ch := range spec {
		switch byte(ch) {
		case lowCh:
			binaryStr.WriteString("0")
		case highCh:
			binaryStr.WriteString("1")
		}
	}

	num, _ := strconv.ParseInt(binaryStr.String(), 2, 64)

	return int(num)
}

func ParseInput(input string) (res []BoardingPass) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, BoardingPass{line})
	}

	return
}

func FindMaxSeatId(passes []BoardingPass) (res int) {
	for _, p := range passes {
		if seatId := p.GetSeatId(); seatId > res {
			res = seatId
		}
	}

	return
}

func AllSeatIds(passes []BoardingPass) (res []int) {
	for _, p := range passes {
		res = append(res, p.GetSeatId())
	}
	sort.Ints(res)

	return
}

func FindMissingNum(sortedSeatIds []int) int {
	firstId := sortedSeatIds[0]
	for i, p := range sortedSeatIds {
		if p-i > firstId {
			return p - 1
		}
	}

	return 0
}
