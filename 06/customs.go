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

func (g CustomGroup) CountEveryoneYeses() (count int) {
	set := make(map[rune]int)
	for _, input := range g.RawInputs {
		for _, ch := range input {
			set[ch]++
		}
	}
	for _, v := range set {
		if v == len(g.RawInputs) {
			count++
		}
	}

	return
}
