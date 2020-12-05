package main

import (
	// "fmt"
	"regexp"
	"strconv"
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

func (p Passport) IsActuallyValid() (res bool) {
	return p.IsValid() && p.isValidByr() && p.isValidEyr() && p.isValidIyr() &&
		p.isValidHgt() && p.isValidEcl() && p.isValidHcl() && p.isValidPid()
}

func (p Passport) isValidByr() bool {
	num, err := strconv.Atoi(p.Fields["byr"])
	return err == nil && num <= 2002 && num >= 1920
}

func (p Passport) isValidIyr() bool {
	num, err := strconv.Atoi(p.Fields["iyr"])
	return err == nil && num <= 2020 && num >= 2010
}

func (p Passport) isValidEyr() bool {
	num, err := strconv.Atoi(p.Fields["eyr"])
	return err == nil && num <= 2030 && num >= 2020
}

func (p Passport) isValidHgt() bool {
	validHgt := regexp.MustCompile(`^([0-9a-f]{2,3})(cm|in)$`)
	if !validHgt.MatchString(p.Fields["hgt"]) {
		return false
	}
	matches := validHgt.FindSubmatch([]byte(p.Fields["hgt"]))
	height, _ := strconv.Atoi(string(matches[1]))
	unit := string(matches[2])
	if unit == "cm" {
		return height >= 150 && height <= 193
	} else { // in
		return height >= 59 && height <= 76
	}
}

func (p Passport) isValidHcl() bool {
	validHcl := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return validHcl.MatchString(p.Fields["hcl"])
}

func (p Passport) isValidEcl() bool {
	validHcl := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return validHcl.MatchString(p.Fields["ecl"])
}

func (p Passport) isValidPid() bool {
	validHcl := regexp.MustCompile(`^[0-9]{9}$`)
	return validHcl.MatchString(p.Fields["pid"])
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
		if check(p) {
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
