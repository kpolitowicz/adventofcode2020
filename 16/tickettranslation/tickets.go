package tickettranslation

type Ticket []int

func TicketScanningErrorRate(tickets []Ticket, rules Rules) (sum int) {
	for _, ticket := range tickets {
		for _, num := range ticket {
			if !validForAnyField(num, rules) {
				sum += num
			}
		}
	}
	return
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
