package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/dockingdata"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	program := dockingdata.ParseInput(strings.Trim(string(dat), "\n"))
	computer := dockingdata.NewComputer()
	computer.Execute(&program)

	switch os.Args[1] {
	case "1":
		sum := uint64(0)
		for _, val := range computer.GetMemory() {
			sum += val
		}
		fmt.Println(sum)
	case "2":
		// variants := jolts.CountChains()
		// fmt.Println(variants)
	}
}
