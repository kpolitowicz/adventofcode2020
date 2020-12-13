package adapters

import (
	// "regexp"
	"sort"
	"strconv"
	"strings"
)

type JoltageAdapters []int

func (ja JoltageAdapters) CountJoltageDiffs() map[int]int {
	res := make(map[int]int)
	sort.Ints(ja)

	prev := 0
	for _, next := range ja {
		diff := next - prev
		res[diff]++
		prev = next
	}

	// add diff between last adapter and the device (always 3)
	res[3]++

	return res
}

func ParseInput(input string) (res JoltageAdapters) {
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		res = append(res, num)
	}

	return
}
