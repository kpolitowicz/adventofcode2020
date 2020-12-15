package dockingdata

import (
	"fmt"
	"math"
	"strconv"
)

var _ = fmt.Println

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

func (c *Computer) AllFloating(num uint64) (res []uint64) {
	strMask, indexesOfXses := c.bitMaskFor(num)
	combinations := c.allFloatingCombinations(len(indexesOfXses))
	for _, variant := range combinations {
		str := ""
		currentX := 0
		for _, ch := range strMask {
			if ch == 'X' {
				str += string(variant[currentX])
				currentX++
			} else {
				str += string(ch)
			}
		}
		variantNum, _ := strconv.ParseUint(str, 2, 64)
		res = append(res, variantNum)
	}

	return
}

func (c *Computer) bitMaskFor(num uint64) (res string, idxX []int) {
	strMask := fmt.Sprintf("%036s", strconv.FormatUint(num, 2))
	for idx, ch := range c.CurrentMask {
		switch ch {
		case '0':
			res += string(strMask[idx])
		case '1':
			res += "1"
		case 'X':
			res += "X"
			idxX = append(idxX, idx)
		}
	}

	return
}

func (c *Computer) allFloatingCombinations(num int) (res []string) {
	for n := uint64(0); float64(n) < math.Pow(2, float64(num)); n++ {
		numStr := strconv.FormatInt(int64(num), 10)
		nStr := fmt.Sprintf("%0"+numStr+"s", strconv.FormatUint(n, 2))
		res = append(res, nStr)
	}

	return
}
