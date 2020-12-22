package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kpolitowicz/adventofcode2020/conwaycubes"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	pu := conwaycubes.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		for i := 0; i < 6; i++ {
			pu = pu.CycleOnce()
		}
		fmt.Println(pu.CountActive())
	case "2":
		hu := conwaycubes.NewHyperUniverse(pu)
		for i := 0; i < 6; i++ {
			hu = hu.CycleOnce()
		}
		fmt.Println(hu.CountActive())
	}
}
