package shuttlesearch

import (
	"fmt"
	"strconv"
	"strings"
)

var _ = fmt.Println

func ParseInput(input string) (res BusSchedule) {
	for idx, line := range strings.Split(input, "\n") {
		if idx == 0 {
			continue //skip 1st line
		}
		res = parseInputLine(line)
	}

	return
}

func parseInputLine(line string) (res BusSchedule) {
	var num int
	for _, numStr := range strings.Split(line, ",") {
		if numStr == "x" {
			num = 0
		} else {
			num, _ = strconv.Atoi(numStr)
		}
		res = append(res, num)
	}
	return
}
