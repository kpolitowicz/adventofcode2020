package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/encoding"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	data := encoding.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		fmt.Println(data.FindBadNumber(25))
	case "2":
		// program.FixCorruptCommand()
		// fmt.Println(program.Accumulator)
	}
}
