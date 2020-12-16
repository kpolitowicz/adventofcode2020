package main

import (
	"fmt"
	"github.com/kpolitowicz/adventofcode2020/tickettranslation"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dat, _ := ioutil.ReadFile("input.txt")
	rules, myTicket, nearbyTickets := tickettranslation.ParseInput(strings.Trim(string(dat), "\n"))

	switch os.Args[1] {
	case "1":
		fmt.Println(tickettranslation.TicketScanningErrorRate(nearbyTickets, rules))
	case "2":
		validTickets := tickettranslation.AllValidTickets(rules, nearbyTickets, myTicket)
		mapping := tickettranslation.TicketFieldsMapping(rules, validTickets)
		myTicket.Print(mapping)
	}
}
