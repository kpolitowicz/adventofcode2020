package tickettranslation

type Ticket []int

func (t Ticket) errorRate(rules Rules) (sum int) {
	for _, num := range t {
		if !validForAnyField(num, rules) {
			sum += num
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
