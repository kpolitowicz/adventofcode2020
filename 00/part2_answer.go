package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	ary := strings.Split(string(dat), "\n")

	fmt.Println(CalcSignalDistance(ary[0], ary[1]))
}
