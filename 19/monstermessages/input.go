package monstermessages

import (
	"sort"
	"strconv"
	"strings"
)

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
