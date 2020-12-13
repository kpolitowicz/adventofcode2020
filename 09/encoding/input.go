package encoding

import (
	// "regexp"
	"strconv"
	"strings"
)

type XmasData []uint64

func ParseInput(input string) (res XmasData) {
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.ParseUint(line, 10, 64)
		res = append(res, num)
	}

	return
}
