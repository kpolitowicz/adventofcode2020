package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/kpolitowicz/adventofcode2020/monstermessages"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	rules, messages := monstermessages.ParseInput(strings.Trim(string(dat), "\n"))
	ruleStr := rules.ResolveRule("0")

	switch os.Args[1] {
	case "1":
		fmt.Println(monstermessages.ValidCount(messages, ruleStr))
	case "2":
		fmt.Println(rules["8"])
		fmt.Println(rules["11"])
		fmt.Println(rules.ResolveRule("42"))
		fmt.Println(rules.ResolveRule("31"))

		testRep := regexp.MustCompile("^(a(?1)?b)$")
		fmt.Println(testRep.MatchString("aabb"))
	}
}
