package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/seating"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	seatMap := seating.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		finalSeating := seatMap.Simulate(seating.FerrySeating.TransformRow)
		fmt.Println(finalSeating.CountOccupied())
	case "2":
		// variants := jolts.CountChains()
		// fmt.Println(variants)
	}
}
