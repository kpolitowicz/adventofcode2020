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
		// validTickets := tickettranslation.AllValidTickets(rules, nearbyTickets, myTicket)
		// mapping := tickettranslation.TicketFieldsMapping(rules, validTickets)
		// mult := 1
		// for _, num := range myTicket.DepartureValues(mapping) {
		// 	mult *= num
		// }
		// fmt.Println(mult)
	}
}
