package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")

	validCheckFun := Password.isValid
	if part := os.Args[1]; part == "2" {
		validCheckFun = Password.isReallyValid
	}

	passwords := ParseInput(strings.Trim(string(dat), "\n"))
	countValid := CountValidPasswords(passwords, validCheckFun)

	fmt.Println(countValid)
}
