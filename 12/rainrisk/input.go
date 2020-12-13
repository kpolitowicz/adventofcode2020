package rainrisk

import (
	"fmt"
	"strconv"
	"strings"
)

var _ = fmt.Println

type NavigationCmd struct {
	Cmd byte
	Arg int
}
type NavigationData []NavigationCmd

func ParseInput(input string) (res NavigationData) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, parseInputLine(line))
	}

	return
}

func parseInputLine(line string) NavigationCmd {
	argNum, _ := strconv.Atoi(line[1:])
	return NavigationCmd{line[0], argNum}
}
