package main

import (
	"sort"
	"strconv"
	"strings"
)

func ConvertInputToListOfNumbers(input string) (res []int) {
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		res = append(res, num)
	}
	return
}

func FindEntries(expenseReport []int) []int {
	minus2020Report := minus2020(expenseReport)
	elem := firstCommonElem(expenseReport, minus2020Report)

	return []int{2020 - elem, elem}
}

func minus2020(ary []int) (res []int) {
	for _, v := range ary {
		res = append(res, 2020-v)
	}

	return
}

func firstCommonElem(ary1, ary2 []int) (res int) {
	sort.Ints(ary1)
	sort.Ints(ary2)

	for _, v1 := range ary1 {
		for _, v2 := range ary2 {
			if v1 == v2 {
				res = v1
				break
			}
		}
	}

	return
}
