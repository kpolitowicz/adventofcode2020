package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kpolitowicz/adventofcode2020/operationorder"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	expressions := operationorder.ParseInput(strings.Trim(string(dat), "\n"))

	var sum uint64 = 0
	switch os.Args[1] {
	case "1":
		for _, expr := range expressions {
			sum += expr.Calculate()
		}
	case "2":
		for _, expr := range expressions {
			sum += expr.AdvCalculate()
		}
	}
	fmt.Println(sum)
}
