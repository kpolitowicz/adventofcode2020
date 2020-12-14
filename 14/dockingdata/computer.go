package dockingdata

import (
	"strconv"
)

type Computer struct {
	CurrentMask string
	Memory      MemMap
	Version     int
}

type MemMap map[uint64]uint64

func NewComputer(version int) *Computer {
	return &Computer{
		CurrentMask: "",
		Memory:      make(MemMap),
		Version:     version,
	}
}

func (c *Computer) GetMask() string {
	return c.CurrentMask
}

func (c *Computer) SetMask(mask string) {
	c.CurrentMask = mask
}

func (c *Computer) GetMemory() MemMap {
	return c.Memory
}

func (c *Computer) SetMemoryAt(addr, value uint64) {
	c.Memory[addr] = value
}

func (c *Computer) Execute(program *Program) {
	for _, cmd := range *program {
		cmd.ExecuteOn(c)
	}
}

func (c *Computer) OrMask() uint64 {
	strMask := ""
	for _, ch := range c.CurrentMask {
		if ch == '1' {
			strMask += "1"
		} else {
			strMask += "0"
		}
	}

	intMask, _ := strconv.ParseUint(strMask, 2, 64)
	return intMask
}
func (c *Computer) AndMask() uint64 {
	strMask := ""
	for _, ch := range c.CurrentMask {
		if ch == '0' {
			strMask += "0"
		} else {
			strMask += "1"
		}
	}

	intMask, _ := strconv.ParseUint(strMask, 2, 64)
	return intMask
}
