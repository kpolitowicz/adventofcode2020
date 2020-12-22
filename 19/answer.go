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
	validMsg := regexp.MustCompile("^" + ruleStr + "$")

	switch os.Args[1] {
	case "1":
		count := 0
		for _, msg := range messages {
			if msg.IsValid(validMsg) {
				count++
			}
		}
		fmt.Println(count)
	case "2":
	}
}
