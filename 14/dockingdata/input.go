package dockingdata

import (
	"fmt"
	// "sort"
	"strconv"
	"strings"
)

var _ = fmt.Println

type Cmd interface {
}

type MaskCmd struct {
	Mask string
}

type MemWriteCmd struct {
	Addr, Value uint
}

type Program []Cmd

func ParseInput(input string) (res Program) {
	var parsedCmd Cmd
	for _, line := range strings.Split(input, "\n") {
		if line[:4] == "mask" {
			parsedCmd = parseMask(line)
		} else {
			parsedCmd = parseMemWriteCmd(line)
		}
		res = append(res, parsedCmd)
	}

	return
}

func parseMask(line string) MaskCmd {
	splits := strings.Split(line, " = ")

	return MaskCmd{splits[1]}
}

func parseMemWriteCmd(line string) MemWriteCmd {
	splits := strings.Split(line, "] = ")
	addr, _ := strconv.ParseUint(splits[0][4:], 10, 32)
	value, _ := strconv.ParseUint(splits[1], 10, 32)

	return MemWriteCmd{uint(addr), uint(value)}
}
