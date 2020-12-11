package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/bags"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	bm := bags.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		visited := make(bags.VisitedMap)
		bm.SearchContainers(bm.GetNode("shiny gold"), &visited)
		fmt.Println(len(visited) - 1)
	case "2":
		counted := make(bags.CountedMap)
		fmt.Println(bm.CountBags(bm.GetNode("shiny gold"), &counted) - 1)
	}
}
