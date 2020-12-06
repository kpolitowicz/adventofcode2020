package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	groups := ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		fmt.Println(SumAllQuestions(groups, CustomGroup.CountAnyoneYeses))
	case "2":
		fmt.Println(SumAllQuestions(groups, CustomGroup.CountEveryoneYeses))
	}
}
