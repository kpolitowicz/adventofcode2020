package tickettranslation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var _ = fmt.Println

type Range struct {
	Min, Max int
}
type Rules map[string]([]Range)

func ParseInput(input string) (Rules, Ticket, []Ticket) {
	var myTicket Ticket
	var nearbyTickets []Ticket
	rules := make(Rules)

	currentInputSection := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			currentInputSection++
			continue
		}
		if line == "your ticket:" || line == "nearby tickets:" {
			continue
		}
		switch currentInputSection {
		case 0: // rules
			label, range1, range2 := parseRulesLine(line)
			rules[label] = []Range{range1, range2}
		case 1: // my ticket
			myTicket = parseTicketLine(line)
		case 2: // nearby tickets
			nearbyTickets = append(nearbyTickets, parseTicketLine(line))
		}
	}

	return rules, myTicket, nearbyTickets
}

func parseRulesLine(line string) (string, Range, Range) {
	validRule := regexp.MustCompile(`^(.*): (\d+)-(\d+) or (\d+)-(\d+)$`)
	matches := validRule.FindSubmatch([]byte(line))

	label := string(matches[1])
	range1Min, _ := strconv.Atoi(string(matches[2]))
	range1Max, _ := strconv.Atoi(string(matches[3]))
	range2Min, _ := strconv.Atoi(string(matches[4]))
	range2Max, _ := strconv.Atoi(string(matches[5]))
	return label, Range{range1Min, range1Max}, Range{range2Min, range2Max}
}

func parseTicketLine(line string) (res Ticket) {
	for _, numStr := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(numStr)
		res = append(res, num)
	}

	return
}
