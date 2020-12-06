package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/customs"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	groups := customs.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		fmt.Println(customs.SumAllQuestions(groups, customs.CustomGroup.CountAnyoneYeses))
	case "2":
		fmt.Println(customs.SumAllQuestions(groups, customs.CustomGroup.CountEveryoneYeses))
	}
}
