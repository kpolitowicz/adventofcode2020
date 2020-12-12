package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/handheld"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	program := handheld.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		program.Execute()
		fmt.Println(program.Accumulator)
	case "2":
		// counted := make(handheld.CountedMap)
		// fmt.Println(bm.CountBags(bm.GetNode("shiny gold"), &counted) - 1)
	}
}
