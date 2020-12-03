package main

import (
	// "fmt"
	"strconv"
	"strings"
)

type Password struct {
	password string
	letter   byte
	min, max int
}

func (p Password) isValid() bool {
	count := 0
	for _, ch := range p.password {
		if byte(ch) == p.letter {
			count++
		}
	}

	return count >= p.min && count <= p.max
}

func ParseInput(input string) (res []Password) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, parseLineIntoPasswordSpec(line))
	}

	return
}

func CountValidPasswords(passwords []Password, validCheck func(Password) bool) int {
	count := 0
	for _, password := range passwords {
		if validCheck(password) {
			count++
		}
	}

	return count
}

// Example of line:
// 1-3 a: abcde
func parseLineIntoPasswordSpec(line string) Password {
	elems := strings.Split(line, " ")
	// fmt.Println(elems)

	allowed_range := strings.Split(elems[0], "-")
	min, _ := strconv.Atoi(allowed_range[0])
	max, _ := strconv.Atoi(allowed_range[1])
	letter_spec := elems[1]
	pass := elems[2]

	return Password{pass, letter_spec[0], min, max}
}
