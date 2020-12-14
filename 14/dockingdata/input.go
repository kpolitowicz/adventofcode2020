package dockingdata

import (
	"fmt"
	// "sort"
	"strconv"
	"strings"
)

var _ = fmt.Println

type Cmd interface {
	ExecuteOn(*Computer)
}

type MaskCmd struct {
	Mask string
}

func (m MaskCmd) ExecuteOn(c *Computer) {
	c.SetMask(m.Mask)
}

type MemWriteCmd struct {
	Addr, Value uint64
}

func (mw MemWriteCmd) ExecuteOn(c *Computer) {
	value := (mw.Value | c.OrMask()) & c.AndMask()
	c.SetMemoryAt(mw.Addr, value)
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
	addr, _ := strconv.ParseUint(splits[0][4:], 10, 64)
	value, _ := strconv.ParseUint(splits[1], 10, 64)

	return MemWriteCmd{addr, value}
}
