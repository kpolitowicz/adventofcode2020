package main

import (
	"fmt"
	"io/ioutil"
	// "os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	slope := ParseInput(strings.Trim(string(dat), "\n"))

	// validCheckFun := Password.isValid
	// if part := os.Args[1]; part == "2" {
	// 	validCheckFun = Password.isReallyValid
	// }

	trees := slope.CountTrees()

	fmt.Println(trees)
}
