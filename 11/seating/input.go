package seating

import (
	// "fmt"
	// "sort"
	// "strconv"
	"strings"
)

type FerrySeating []string

func ParseInput(input string) (res FerrySeating) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, line)
	}

	return
}
