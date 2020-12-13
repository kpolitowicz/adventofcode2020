package encoding

import (
	// "regexp"
	"strconv"
	"strings"
)

type XmasData []uint64

func (d XmasData) FindBadNumber(preambleSize int) (res uint64) {
	for i := preambleSize; i < len(d); i++ {
		if !d.goodNumber(i, preambleSize) {
			return d[i]
		}
	}

	return
}

func (d XmasData) goodNumber(index, preambleSize int) bool {
	num := d[index]

	for i := index - preambleSize; i < index-1; i++ {
		for j := i + 1; j < index; j++ {
			if num == d[i]+d[j] {
				return true
			}
		}
	}

	return false
}

func ParseInput(input string) (res XmasData) {
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.ParseUint(line, 10, 64)
		res = append(res, num)
	}

	return
}
