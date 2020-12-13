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
	badNum := data.FindBadNumber(25)

	switch os.Args[1] {
	case "1":
		fmt.Println(badNum)
	case "2":
		low, high := data.FindSetForWeakness(badNum)
		min := data[low]
		max := data[high]
		for i := low + 1; i <= high; i++ {
			if data[i] < min {
				min = data[i]
			}
			if data[i] > max {
				max = data[i]
			}
		}
		fmt.Println(min + max)
	}
}
