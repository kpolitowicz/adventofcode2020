package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/shuttlesearch"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	schedule := shuttlesearch.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		// this is also answer to part 2
		// way faster than teh other case
		// calculation based on the fact that matching_timestamp + 13 will be divisible by:
		// 13, 659, 19, 23, 29 (so by 108569591)
		for i := uint64(108569591); ; i += uint64(108569591) {
			if (i-10)%41 == 0 && (i-6)%37 == 0 && (i+31)%409 == 0 && (i-3)%17 == 0 {
				fmt.Println(i - 13)
				break
			}
		}
	case "2":
		fmt.Println(schedule)
		fmt.Println(schedule.FindEarliestTimestamp())
	}
}
