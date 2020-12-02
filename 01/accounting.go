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

func FindEntries3(expenseReport []int) (res []int) {
	sort.Ints(expenseReport)

	idx1 := 0
	idx2 := 1
	idx3 := 2
	reportLen := len(expenseReport)

	for idx3 < reportLen {
		num1 := expenseReport[idx1]
		num2 := expenseReport[idx2]
		num3 := expenseReport[idx3]
		sum := num1 + num2 + num3

		if sum == 2020 {
			break
		}
		if sum < 2020 {
			if idx3 < reportLen-1 {
				idx3++
				continue
			}
			if idx2 < reportLen-2 {
				idx2++
				idx3 = idx2 + 1
				continue
			}
			if idx1 < reportLen-3 {
				idx1++
				idx2 = idx1 + 1
				idx3 = idx1 + 2
				continue
			}
		}
		if sum > 2020 {
			if idx2 < reportLen-2 {
				idx2++
				idx3 = idx2 + 1
				continue
			}
			if idx1 < reportLen-3 {
				idx1++
				idx2 = idx1 + 1
				idx3 = idx1 + 2
				continue
			}
		}
	}

	return []int{
		expenseReport[idx1],
		expenseReport[idx2],
		expenseReport[idx3],
	}
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
