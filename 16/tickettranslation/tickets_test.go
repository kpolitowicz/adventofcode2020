package tickettranslation

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestTicketScanningErrorRate(t *testing.T) {
	rules, _, nearbyTickets := ParseInput(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	got := TicketScanningErrorRate(nearbyTickets, rules)
	want := 71

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAllValidTickets(t *testing.T) {
	rules, myTicket, nearbyTickets := ParseInput(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	got := AllValidTickets(rules, nearbyTickets, myTicket)
	want := []Ticket{
		{7, 1, 14},
		{7, 3, 47},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTicketFieldsMapping(t *testing.T) {
	rules := Rules{
		"class": []Range{{0, 1}, {4, 19}},
		"row":   []Range{{0, 5}, {8, 19}},
		"seat":  []Range{{0, 13}, {16, 19}},
	}
	tickets := []Ticket{
		{11, 12, 13},
		{3, 9, 18},
		{15, 1, 5},
		{5, 14, 9},
	}

	got := TicketFieldsMapping(rules, tickets)
	want := TicketFieldsMaybeMap{
		"class": []int{1},
		"row":   []int{0},
		"seat":  []int{2},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
