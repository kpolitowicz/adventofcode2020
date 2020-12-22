package operationorder

import (
	"strings"
)

func ParseInput(input string) (res []mathExpression) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, mathExpression(line))
	}
	return
}
