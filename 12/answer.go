package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/rainrisk"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	naviData := rainrisk.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		ferry := rainrisk.NewFerry()
		ferry.ExecuteNavigation(naviData)
		fmt.Println(math.Abs(float64(ferry.Pos.X)) + math.Abs(float64(ferry.Pos.Y)))
	case "2":
		// variants := jolts.CountChains()
		// fmt.Println(variants)
	}
}
