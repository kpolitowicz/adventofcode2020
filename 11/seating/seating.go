package seating

import "fmt"

var _ = fmt.Println

type FerrySeating []string
type BehaviorModel func(FerrySeating, int, string) string

func (fs FerrySeating) String() (res string) {
	res = "\n"
	for _, row := range fs {
		res += row + "\n"
	}
	return
}

func (fs FerrySeating) Simulate(behaviorModel BehaviorModel) (res FerrySeating) {
	old := fs
	for {
		res = old.NextRound(behaviorModel)
		if old.equal(res) {
			break
		} else {
			old = res
		}
	}

	return
}

func (fs FerrySeating) NextRound(behaviorModel BehaviorModel) (res FerrySeating) {
	for idx, row := range fs {
		res = append(res, behaviorModel(fs, idx, row))
	}

	return
}

func (fs FerrySeating) TransformRow(rowIdx int, row string) (res string) {
	for colIdx, seat := range row {
		switch seat {
		case '.':
			// floor stays floor
			res += "."
		case 'L':
			// change to # if no adjacent occupied
			if fs.countOccupiedAdj(rowIdx, colIdx) == 0 {
				res += "#"
			} else {
				res += "L"
			}
		case '#':
			// change to L if 4+ adjacent occupied
			if fs.countOccupiedAdj(rowIdx, colIdx) >= 4 {
				res += "L"
			} else {
				res += "#"
			}
		}
	}

	return
}

func (fs FerrySeating) TransformRowBetterModel(rowIdx int, row string) (res string) {
	for colIdx, seat := range row {
		switch seat {
		case '.':
			// floor stays floor
			res += "."
		case 'L':
			// change to # if no adjacent occupied
			if fs.countOccupiedInSight(rowIdx, colIdx) == 0 {
				res += "#"
			} else {
				res += "L"
			}
		case '#':
			// change to L if 4+ adjacent occupied
			if fs.countOccupiedInSight(rowIdx, colIdx) >= 5 {
				res += "L"
			} else {
				res += "#"
			}
		}
	}

	fmt.Println(row)
	fmt.Println(res)
	return
}

func (fs FerrySeating) countOccupiedAdj(row, col int) (count int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if fs.isOccupied(row+i, col+j) {
				count++
			}
		}
	}

	return
}

func (fs FerrySeating) countOccupiedInSight(row, col int) (count int) {
	// check up
	for i := row - 1; i >= 0; i-- {
		status := fs.isSeatAndOccupied(i, col)
		if status >= 0 {
			count += status
			break
		}
	}
	// check down
	for i := row + 1; i < len(fs); i++ {
		status := fs.isSeatAndOccupied(i, col)
		if status >= 0 {
			count += status
			break
		}
	}
	// check left
	for i := col - 1; i >= 0; i-- {
		status := fs.isSeatAndOccupied(row, i)
		if status >= 0 {
			count += status
			break
		}
	}
	// check right
	for i := col + 1; i < len(fs[0]); i++ {
		status := fs.isSeatAndOccupied(row, i)
		if status >= 0 {
			count += status
			break
		}
	}
	// check up-left
	for i := row - 1; i >= 0; i-- {
		status := fs.isSeatAndOccupied(i, col-(row-i))
		if status >= 0 {
			count += status
			break
		}
	}
	// check up-right
	for i := row - 1; i >= 0; i-- {
		status := fs.isSeatAndOccupied(i, col+(row-i))
		if status >= 0 {
			count += status
			break
		}
	}
	// check down-right
	for i := row + 1; i < len(fs); i++ {
		status := fs.isSeatAndOccupied(i, col+(i-row))
		if status >= 0 {
			count += status
			break
		}
	}
	// check down-left
	for i := row + 1; i < len(fs); i++ {
		status := fs.isSeatAndOccupied(i, col-(i-row))
		if status >= 0 {
			count += status
			break
		}
	}

	return
}

func (fs FerrySeating) isOccupied(row, col int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(fs) || col >= len(fs[0]) {
		return false
	}
	if row < 0 || col < 0 {
		return false
	}
	return fs[row][col] == '#'
}

// returns:
// 1 if it is a seat and is occupied
// 0 if it is an empty seat
// -1 if floor
func (fs FerrySeating) isSeatAndOccupied(row, col int) int {
	if row < 0 || col < 0 {
		return -1
	}
	if row >= len(fs) || col >= len(fs[0]) {
		return -1
	}
	switch fs[row][col] {
	case '#':
		return 1
	case 'L':
		return 0
	}
	return -1
}

func (fs FerrySeating) equal(anotherFs FerrySeating) bool {
	for rowIdx, row := range fs {
		for colIdx, seat := range row {
			if byte(seat) != anotherFs[rowIdx][colIdx] {
				return false
			}
		}
	}

	return true
}

func (fs FerrySeating) CountOccupied() (count int) {
	for _, row := range fs {
		for _, seat := range row {
			if seat == '#' {
				count++
			}
		}
	}

	return
}
