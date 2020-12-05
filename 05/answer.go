package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	passes := ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		fmt.Println(FindMaxSeatId(passes))
	case "2":
		// validityCheck = Passport.IsActuallyValid
	}
}
