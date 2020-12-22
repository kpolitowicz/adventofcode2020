package conwaycubes

import (
	"fmt"
	"strconv"
	"strings"
)

var _ = fmt.Println

type pocketUniverse []zLayer

func (pu pocketUniverse) String() (res string) {
	for z, layer := range pu {
		res += "z=" + strconv.FormatInt(int64(z), 10) + "\n"
		res += fmt.Sprintf("%v\n", layer)
	}
	return
}

type zLayer []string

func (l zLayer) String() (res string) {
	for _, str := range l {
		res += fmt.Sprintf("%s\n", str)
	}
	return
}

// ParseInput converts the content of `input.txt`
// into the problem structure.
func ParseInput(input string) pocketUniverse {
	layer := zLayer{}
	for _, line := range strings.Split(input, "\n") {
		layer = append(layer, line)
	}
	return pocketUniverse{layer}
}
