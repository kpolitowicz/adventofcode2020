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

	switch os.Args[1] {
	case "1":
		var sum uint64 = 0
		for _, expr := range expressions {
			sum += expr.Calculate()
		}
		fmt.Println(sum)
	case "2":
	}
}
