package adapters

import (
	// "fmt"
	"sort"
	"strconv"
	"strings"
)

type JoltageAdapters []int

func (ja JoltageAdapters) CountJoltageDiffs() map[int]int {
	res := make(map[int]int)

	prev := 0
	for _, next := range ja {
		if next == 0 {
			continue
		}
		diff := next - prev
		res[diff]++
		prev = next
	}

	return res
}

func ParseInput(input string) (res JoltageAdapters) {
	res = JoltageAdapters{0}
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		res = append(res, num)
	}

	sort.Ints(res)
	res = append(res, res[len(res)-1]+3) //add device as +3

	return
}
