package handheld

import (
	// "regexp"
	"strconv"
	"strings"
)

func ParseInput(input string) (res Program) {
	for _, line := range strings.Split(input, "\n") {
		res.Commands = append(res.Commands, parseInputLine(line))
	}

	return
}

func parseInputLine(line string) Command {
	parsedCommand := strings.Split(line, " ")

	return Command{
		Instruction: parsedCommand[0],
		Arg:         parseArg(parsedCommand[1]),
	}
}

func parseArg(arg string) (res int) {
	res, _ = strconv.Atoi(arg[1:])
	if arg[0] == '-' {
		res = -res
	}

	return
}
