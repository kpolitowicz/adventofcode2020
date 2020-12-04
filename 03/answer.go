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

	switch os.Args[1] {
	case "1":
		trees := slope.CountTrees(3, 1)
		fmt.Println(trees)
	case "2":
		product :=
			slope.CountTrees(1, 1) *
				slope.CountTrees(3, 1) *
				slope.CountTrees(5, 1) *
				slope.CountTrees(7, 1) *
				slope.CountTrees(1, 2)
		fmt.Println(product)
	}
}
