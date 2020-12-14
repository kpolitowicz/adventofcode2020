package shuttlesearch

import "fmt"

var _ = fmt.Println

type BusSchedule []int

func (s BusSchedule) FindEarliestTimestamp() uint64 {
	maxBusId := 0
	maxBusIdIdx := 0
	for idx, busId := range s {
		if busId > maxBusId {
			maxBusId = busId
			maxBusIdIdx = idx
		}
	}

	for i := uint64(maxBusId); ; i += uint64(maxBusId) {
		ts := i - uint64(maxBusIdIdx)
		if s.checkAtTimestamp(ts) {
			return ts
		}
	}

	return 0
}

func (s BusSchedule) checkAtTimestamp(ts uint64) bool {
	for idx, busId := range s {
		if busId == 0 {
			continue
		}
		if int((ts+uint64(idx))%uint64(busId)) != 0 {
			return false
		}
	}

	return true
}
