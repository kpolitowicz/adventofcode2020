package main

import (
	// "fmt"
	"strings"
)

type PassportFields map[string]string
type Passport struct {
	Fields PassportFields
}

func (p Passport) IsValid() (res bool) {
	res = true

	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	for _, key := range requiredFields {
		_, ok := p.Fields[key]
		res = res && ok
	}

	return
}

func ParseInput(input string) (res []Passport) {
	currentPassportFields := make(PassportFields)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			res = append(res, Passport{currentPassportFields})
			currentPassportFields = make(PassportFields)
		} else {
			parseLineIntoFields(line, &currentPassportFields)
		}
	}

	res = append(res, Passport{currentPassportFields})
	return
}

func CountValid(passports []Passport, check func(Passport) bool) (count int) {
	for _, p := range passports {
		if p.IsValid() {
			count++
		}
	}

	return
}

func parseLineIntoFields(line string, passportFields *PassportFields) {
	for _, fieldSpec := range strings.Split(line, " ") {
		parsedKeyVal := strings.Split(fieldSpec, ":")
		(*passportFields)[parsedKeyVal[0]] = parsedKeyVal[1]
	}
}
