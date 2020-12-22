package conwaycubes

import (
	"fmt"
	"strings"
)

var _ = fmt.Println

// ParseInput converts the content of `input.txt`
// into the problem structure.
func ParseInput(input string) pocketUniverse {
	layer := zLayer{}
	for _, line := range strings.Split(input, "\n") {
		layer = append(layer, line)
	}
	return pocketUniverse{layer}
}
