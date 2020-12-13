package adapters

import (
	// "regexp"
	// "sort"
	"strconv"
	"strings"
)

type JoltageAdapters []int

func ParseInput(input string) (res JoltageAdapters) {
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		res = append(res, num)
	}

	return
}
