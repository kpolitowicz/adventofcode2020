package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	passports := ParseInput(strings.Trim(string(dat), "\n"))

	var validityCheck func(Passport) bool
	switch os.Args[1] {
	case "1":
		validityCheck = Passport.IsValid
	case "2":
		validityCheck = Passport.IsActuallyValid
	}

	fmt.Println(CountValid(passports, validityCheck))
}
