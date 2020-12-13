package seating

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var _ = fmt.Println

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

func (ja JoltageAdapters) CountChains() uint64 {
	res := uint64(1)

	prev := 0
	currentStreakLen := 0 // streak is the chain of increase-by-1 joltages
	for _, next := range ja {
		diff := next - prev
		prev = next
		if diff > 1 {
			res *= ja.combinationsForStreakOf(currentStreakLen)
			currentStreakLen = 1
		} else {
			currentStreakLen++
		}
	}

	return res
}

func (ja JoltageAdapters) combinationsForStreakOf(length int) uint64 {
	switch length {
	case 0, 1, 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	case 5:
		return 7
	default:
		return 0
	}
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
