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
	ferry := rainrisk.NewFerry()

	switch os.Args[1] {
	case "1":
		ferry.ExecuteNavigation(naviData, (*rainrisk.Ferry).ExecuteCommand)
	case "2":
		ferry.ExecuteNavigation(naviData, (*rainrisk.Ferry).ExecuteWaypointCommand)
	}
	fmt.Println(math.Abs(float64(ferry.Pos.X)) + math.Abs(float64(ferry.Pos.Y)))
}
