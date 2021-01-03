package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kpolitowicz/adventofcode2020/jurassicjigsaw"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	images := jurassicjigsaw.ParseInput(strings.Trim(string(dat), "\n"))
	edges := images.GetEdges()

	switch os.Args[1] {
	case "1":
		for title, edges := range *edges.FindSigletons() {
			fmt.Println(title)
			for _, edge := range edges {
				fmt.Println(edge)
			}
		}
		// Only these 4 tiles have 4 entries (2 sides, flipped and not)
		// Thus, these are the corner tiles:
		// 1213 * 1291 * 1543 * 1117 = 2699020245973
	case "2":
		// aaa
	}
}
