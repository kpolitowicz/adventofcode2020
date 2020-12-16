package tickettranslation

import (
	"fmt"
	"reflect"
	"testing"
)

var _ = fmt.Println

func TestParseInput(t *testing.T) {
	gotRules, gotMyTicket, gotNearbyTickets := ParseInput(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	wantRules := Rules{
		"class": []Range{{1, 3}, {5, 7}},
		"row":   []Range{{6, 11}, {33, 44}},
		"seat":  []Range{{13, 40}, {45, 50}},
	}
	if !reflect.DeepEqual(gotRules, wantRules) {
		t.Errorf("got %v want %v", gotRules, wantRules)
	}

	wantMyTicket := Ticket{7, 1, 14}
	if !reflect.DeepEqual(gotMyTicket, wantMyTicket) {
		t.Errorf("got %v want %v", gotMyTicket, wantMyTicket)
	}

	wantNearbyTickets := []Ticket{
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}
	if !reflect.DeepEqual(gotNearbyTickets, wantNearbyTickets) {
		t.Errorf("got %v want %v", gotNearbyTickets, wantNearbyTickets)
	}
}
