package operationorder

import (
	"strconv"
	"strings"
)

type mathExpression string

func newExpression(s string) mathExpression {
	return mathExpression(s)
}

func (e mathExpression) Calculate() (res uint64) {
	if strings.ContainsRune(string(e), '(') {
		return e.simplifyParen().Calculate()
	}
	currentOp := '+'
	for _, token := range strings.Split(string(e), " ") {
		switch token {
		case "+":
			currentOp = '+'
		case "*":
			currentOp = '*'
		default:
			res = e.calculateNext(res, token, currentOp)
		}

	}

	return res
}

func (e mathExpression) calculateNext(current uint64, numStr string, op rune) uint64 {
	num, _ := strconv.ParseUint(numStr, 10, 64)
	if op == '+' {
		return current + num
	}
	return current * num
}

func (e mathExpression) simplifyParen() mathExpression {
	nesting := 0
	subExpr := ""
	simplifiedExpr := ""
	for _, ch := range e {
		switch ch {
		case '(':
			if nesting > 0 {
				subExpr += string(ch)
			}
			nesting++
		case ')':
			nesting--
			if nesting > 0 {
				subExpr += string(ch)
			} else {
				numStr := strconv.FormatUint(newExpression(subExpr).Calculate(), 10)
				simplifiedExpr += numStr
				subExpr = ""
			}
		default:
			if nesting > 0 {
				subExpr += string(ch)
			} else {
				simplifiedExpr += string(ch)
			}
		}
	}
	return newExpression(simplifiedExpr)
}
