package operationorder

import (
	"fmt"
	"strings"
)

var _ = fmt.Println

type mathExpression string

func ParseInput(input string) (res []mathExpression) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, mathExpression(line))
	}
	return
}
