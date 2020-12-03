package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	passwords := ParseInput(strings.Trim(string(dat), "\n"))
	countValid := CountValidPasswords(passwords, Password.isValid)

	fmt.Println(countValid)
}
