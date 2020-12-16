package tickettranslation

import (
	"fmt"
	"strconv"
	"strings"
)

var _ = fmt.Println

type Ticket []int
type TicketFieldsMaybeMap map[string]([]int)

func (t Ticket) errorRate(rules Rules) (sum int) {
	for _, num := range t {
		if !validForAnyField(num, rules) {
			sum += num
		}
	}
	return
}

func (t Ticket) Print(fm TicketFieldsMaybeMap) {
	for label, fieldAry := range fm {
		field := fieldAry[0]
		fmt.Println(label + ": " + strconv.FormatInt(int64(t[field]), 10))
	}
}

func (t Ticket) DepartureValues(fm TicketFieldsMaybeMap) (res []int) {
	for label, fieldAry := range fm {
		if strings.HasPrefix(label, "departure") {
			field := fieldAry[0]
			res = append(res, t[field])
		}
	}
	return
}

func TicketScanningErrorRate(tickets []Ticket, rules Rules) (sum int) {
	for _, ticket := range tickets {
		sum += ticket.errorRate(rules)
	}
	return
}

func AllValidTickets(rules Rules, tickets []Ticket, myTicket Ticket) []Ticket {
	res := []Ticket{myTicket}
	for _, ticket := range tickets {
		if ticket.errorRate(rules) == 0 {
			res = append(res, ticket)
		}
	}
	return res
}

func TicketFieldsMapping(rules Rules, tickets []Ticket) TicketFieldsMaybeMap {
	fieldsMap := make(TicketFieldsMaybeMap)

	fieldsLen := len(tickets[0])
	for idx := 0; idx < fieldsLen; idx++ {
		nums := allNumbersFor(idx, tickets)
		for label, ranges := range rules {
			if allNumsValidFor(nums, ranges) {
				fieldsMap[label] = append(fieldsMap[label], idx)
			}
		}
	}

	return pinpointFields(fieldsMap)
}

func validForAnyField(num int, rules Rules) bool {
	for _, ranges := range rules {
		for _, r := range ranges {
			if num >= r.Min && num <= r.Max {
				return true
			}
		}
	}
	return false
}

func allNumbersFor(idx int, tickets []Ticket) (res []int) {
	for _, t := range tickets {
		res = append(res, t[idx])
	}
	return
}

func allNumsValidFor(nums []int, ranges []Range) bool {
	for _, num := range nums {
		if (num < ranges[0].Min || num > ranges[0].Max) && (num < ranges[1].Min || num > ranges[1].Max) {
			return false
		}
	}
	return true
}

func pinpointFields(old TicketFieldsMaybeMap) TicketFieldsMaybeMap {
	pinpointed := []int{}
	for _, cols := range old {
		if len(cols) == 1 {
			pinpointed = append(pinpointed, cols[0])
		}
	}
	if len(pinpointed) == len(old) {
		return old
	}

	res := make(TicketFieldsMaybeMap)
	for label, cols := range old {
		if len(cols) == 1 {
			res[label] = old[label]
			continue
		}
		for _, idx := range cols {
			idxPinpointed := false
			for _, p := range pinpointed {
				if idx == p {
					idxPinpointed = true
				}
			}
			if !idxPinpointed {
				res[label] = append(res[label], idx)
			}
		}
	}
	return pinpointFields(res)
}
