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
		// product :=
		// 	slope.CountTrees(1, 1) *
		// 		slope.CountTrees(3, 1) *
		// 		slope.CountTrees(5, 1) *
		// 		slope.CountTrees(7, 1) *
		// 		slope.CountTrees(1, 2)
		// fmt.Println(product)
	}

	fmt.Println(CountValid(passports, validityCheck))
}
