package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kpolitowicz/adventofcode2020/monstermessages"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	rules, messages := monstermessages.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		ruleStr := rules.ResolveRule("0")
		fmt.Println(monstermessages.ValidCount(messages, ruleStr))
	case "2":
		rules["8"] = "42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42 42 | 42 42 42 42 42 42"
		rules["11"] = "42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31"
		ruleStr := rules.ResolveRule("0")
		fmt.Println(monstermessages.ValidCount(messages, ruleStr))
	}
}
