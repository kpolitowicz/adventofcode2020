package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	slope := ParseInput(strings.Trim(string(dat), "\n"))

	if part := os.Args[1]; part == "1" {
		trees := slope.CountTrees(3, 1)
		fmt.Println(trees)
	} else {
	}
}
