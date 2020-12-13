package seating

type FerrySeating []string

func (fs FerrySeating) Simulate() (res FerrySeating) {
	old := fs
	for {
		res = old.NextRound()
		if old.equal(res) {
			break
		} else {
			old = res
		}
	}

	return
}

func (fs FerrySeating) NextRound() (res FerrySeating) {
	for idx, row := range fs {
		res = append(res, fs.transformRow(idx, row))
	}

	return
}

func (fs FerrySeating) transformRow(rowIdx int, row string) (res string) {
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
		if fs.isOccupied(i, col) {
			count++
			break
		}
	}
	// check down
	for i := row + 1; i < len(fs); i++ {
		if fs.isOccupied(i, col) {
			count++
			break
		}
	}
	// check left
	for i := col - 1; i >= 0; i-- {
		if fs.isOccupied(row, i) {
			count++
			break
		}
	}
	// check right
	for i := col + 1; i < len(fs[0]); i++ {
		if fs.isOccupied(row, i) {
			count++
			break
		}
	}
	// check up-left
	for i := row - 1; i >= 0; i-- {
		if fs.isOccupied(i, col-(row-i)) {
			count++
			break
		}
	}
	// check up-right
	for i := row - 1; i >= 0; i-- {
		if fs.isOccupied(i, col+(row-i)) {
			count++
			break
		}
	}
	// check down-right
	for i := row + 1; i < len(fs); i++ {
		if fs.isOccupied(i, col+(i-row)) {
			count++
			break
		}
	}
	// check down-left
	for i := row + 1; i < len(fs); i++ {
		if fs.isOccupied(i, col-(i-row)) {
			count++
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
