package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/adapters"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	jolts := adapters.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		diffs := jolts.CountJoltageDiffs()
		fmt.Println(jolts)
		fmt.Println(diffs[1] * diffs[3])
	case "2":
		variants := jolts.CountChains()
		fmt.Println(variants)
	}
}
