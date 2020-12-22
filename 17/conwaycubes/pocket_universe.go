package conwaycubes

import (
	"fmt"
	"strconv"
)

type zLayer []string

func (l zLayer) String() (res string) {
	for _, str := range l {
		res += fmt.Sprintf("%s\n", str)
	}
	return
}

type pocketUniverse []zLayer

func (pu pocketUniverse) String() (res string) {
	for z, layer := range pu {
		res += "z=" + strconv.FormatInt(int64(z), 10) + "\n"
		res += fmt.Sprintf("%v\n", layer)
	}
	return
}

func (pu pocketUniverse) CountActive() (count int) {
	for z := 0; z < len(pu); z++ {
		for row := 0; row < len(pu[0]); row++ {
			for col := 0; col < len(pu[0][0]); col++ {
				if pu[z][row][col] == '#' {
					count++
				}
			}
		}
	}
	return
}

func (pu pocketUniverse) CycleOnce() (res pocketUniverse) {
	for z := -1; z <= len(pu); z++ {
		layer := zLayer{}
		for row := -1; row <= len(pu[0]); row++ {
			str := ""
			for col := -1; col <= len(pu[0][0]); col++ {
				str += pu.calcNextCycleCube(z, row, col)
			}
			layer = append(layer, str)
		}
		res = append(res, layer)
	}
	return
}

func (pu pocketUniverse) calcNextCycleCube(z, r, c int) string {
	currentCube := pu.currentCube(z, r, c)
	activeNeighbors := pu.countActiveNeighbors(z, r, c)
	switch currentCube {
	case '.':
		if activeNeighbors == 3 {
			return "#"
		} else {
			return "."
		}
	case '#':
		if activeNeighbors == 2 || activeNeighbors == 3 {
			return "#"
		} else {
			return "."
		}
	}
	return "."
}

func (pu pocketUniverse) currentCube(z, r, c int) byte {
	if pu.invalidCoord(z, r, c) {
		return '.'
	}
	return pu[z][r][c]
}

func (pu pocketUniverse) countActiveNeighbors(z, r, c int) (count int) {
	for dz := -1; dz <= 1; dz++ {
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dz == 0 && dr == 0 && dc == 0 {
					continue
				}
				newz := z + dz
				newr := r + dr
				newc := c + dc
				if pu.invalidCoord(newz, newr, newc) {
					continue
				}
				if pu[newz][newr][newc] == '#' {
					count++
				}
			}
		}
	}
	return
}

func (pu pocketUniverse) invalidCoord(z, r, c int) bool {
	if z < 0 || r < 0 || c < 0 {
		return true
	}
	if z >= len(pu) || r >= len(pu[0]) || c >= len(pu[0][0]) {
		return true
	}
	return false
}
