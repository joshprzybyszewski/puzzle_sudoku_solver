package solve

type puzzle struct {
	boxes [16]box
	box   [16][16]uint8

	grid [16][16]value

	cannots [16][16]bits

	placedRows [16]bits
	placedCols [16]bits

	remaining     [16][16]uint8
	remainingRows [16]uint8

	recentlyPlaced bits

	size uint8

	hasEasy bool
}

func NewPuzzle(
	input [][]uint8,
) puzzle {
	switch len(input) {
	case 9:
		return newPuzzle(
			input,
			defaultNineBoxes,
			defaultNineBoxLookups,
		)
	case 12:
		return newPuzzle(
			input,
			defaultTwelveBoxes,
			defaultTwelveBoxLookups,
		)
	case 16:
		return newPuzzle(
			input,
			defaultSixteenBoxes,
			defaultSixteenBoxLookups,
		)
	}

	panic(`dev error`)
}

func NewPuzzleWithBoxLookups(
	input [][]uint8,
	boxNums [][]uint8,
) puzzle {

	var lookup [16][16]uint8
	for i := range boxNums {
		for j := range boxNums[i] {
			lookup[i][j] = uint8(boxNums[i][j] - 1)
		}
	}
	var boxcells [16]box
	size := uint8(len(boxNums))
	for i := range boxNums {
		boxcells[i] = createBox(uint8(i), lookup, size)
	}

	return newPuzzle(
		input,
		boxcells,
		lookup,
	)
}

func newPuzzle(
	input [][]uint8,
	boxes [16]box,
	boxIndexes [16][16]uint8,
) puzzle {
	puzz := puzzle{
		size:  uint8(len(input)),
		boxes: boxes,
		box:   boxIndexes,
	}

	var allCannots bits
	for i := value(puzz.Size() + 1); i <= 16; i++ {
		allCannots |= i.bit()
	}

	for r := uint8(0); r < puzz.Size(); r++ {
		for c := uint8(0); c < puzz.Size(); c++ {
			puzz.cannots[r][c] = allCannots
			puzz.remaining[r][c] = puzz.Size()
		}
		puzz.remainingRows[r] = puzz.Size()
	}

	for r := uint8(0); r < puzz.Size(); r++ {
		for c := uint8(0); c < puzz.Size(); c++ {
			if len(input[r]) == 0 || input[r][c] == 0 {
				continue
			}
			if !puzz.place(r, c, value(input[r][c])) {
				panic(`dev error`)
			}
		}
	}

	if !puzz.FinishInitialPlaces() {
		panic(`dev error`)
	}

	return puzz
}

func (p *puzzle) Size() uint8 {
	return p.size
}

func (p *puzzle) IsSet(r, c uint8) bool {
	return p.grid[r][c] != 0
}

func (p *puzzle) CanPlace(r, c uint8, val value) bool {
	return p.cannots[r][c]&(val.bit()) == 0
}

func (p *puzzle) Val(r, c uint8) value {
	return p.grid[r][c]
}

func (p *puzzle) FinishInitialPlaces() bool {
	for p.hasEasy {
		p.hasEasy = false
		if !p.checkAllForLast() {
			return false
		}

		placed := p.recentlyPlaced
		p.recentlyPlaced = 0

		for v := value(1); v <= value(p.Size()); v++ {
			if placed&(v.bit()) == 0 {
				continue
			}

			if !p.checkBoxEliminations(v) {
				return false
			}

			if !p.validate(v) {
				return false
			}

			placed ^= v.bit()
			if placed == 0 {
				break
			}
		}
	}

	return true
}

func (p *puzzle) placeLast(r, c uint8) bool {
	if p.grid[r][c] != 0 {
		return true
	}
	if p.remaining[r][c] != 1 {
		panic(`dev error`)
	}
	val := value(1)
	for p.cannots[r][c]&(val.bit()) != 0 {
		val++
	}

	return p.place(r, c, val)
}

func (p *puzzle) Place(r, c uint8, val value) bool {
	ok := p.place(r, c, val)
	if !ok {
		return false
	}

	for p.hasEasy {
		p.hasEasy = false
		if !p.checkAllForLast() {
			return false
		}

		placed := p.recentlyPlaced
		p.recentlyPlaced = 0

		for v := value(1); v <= value(p.Size()); v++ {
			if placed&(v.bit()) == 0 {
				continue
			}

			if !p.checkBoxEliminations(v) {
				return false
			}

			if !p.validate(v) {
				return false
			}

			placed ^= v.bit()
			if placed == 0 {
				break
			}
		}
	}

	return true
}

func (p *puzzle) place(r, c uint8, val value) bool {
	if p.grid[r][c] == val {
		return true
	}

	b := val.bit()
	if p.cannots[r][c]&b == b {
		return false
	}

	if val > value(p.Size()) || p.grid[r][c] != 0 {
		panic(`dev error`)
	}

	p.grid[r][c] = val
	p.cannots[r][c] = allBits
	p.remaining[r][c] = 0
	p.placedRows[r] |= b
	p.placedCols[c] |= b
	p.remainingRows[r]--

	for other := uint8(0); other < p.Size(); other++ {
		// update this column (iterate through all the rows)
		if other != r && !p.removeOption(other, c, b) {
			return false
		}
		// update this row (iterate through all the cols)
		if other != c && !p.removeOption(r, other, b) {
			return false
		}
	}

	// update this box (iterate through the 16 nearby)
	bi := p.getBoxIndex(r, c)
	var bc boxCell
	for bci := uint8(0); bci < p.size; bci++ {
		bc = p.boxes[bi][bci]
		if bc.row == r && bc.col == c {
			continue
		}
		if !p.removeOption(bc.row, bc.col, b) {
			return false
		}
	}

	p.recentlyPlaced |= b
	return true
}

func (p *puzzle) removeOption(r, c uint8, b bits) bool {
	if p.cannots[r][c]&b == b {
		return true
	}
	if p.remaining[r][c] == 1 {
		// This removes the last option for this cell.
		// Impossible.
		return false

	}
	p.remaining[r][c]--
	p.cannots[r][c] |= b

	if p.remaining[r][c] == 1 {
		p.hasEasy = true
	}

	return true
}

func (p *puzzle) checkAllForLast() bool {
	var c uint8

	for r := uint8(0); r < p.Size(); r++ {
		for c = uint8(0); c < p.Size(); c++ {
			if p.remaining[r][c] != 1 {
				continue
			}

			if !p.placeLast(r, c) {
				return false
			}
		}
	}

	return true
}

func (p *puzzle) checkBoxEliminations(
	v value,
) bool {
	b := v.bit()

	var r, c uint8
	var last uint8

	var hasBox bool
	var bi, bi2, bci uint8

	for r = 0; r < p.Size(); r++ {
		if p.placedRows[r]&b != 0 {
			continue
		}
		hasBox = false
		last = p.Size()
		for c = 0; c < p.Size(); c++ {
			if p.cannots[r][c]&b != 0 {
				// cannot play it here
				continue
			}

			if last == p.Size() {
				last = c
			} else {
				last = p.Size() + 1
			}

			bi2 = p.getBoxIndex(r, c)
			if !hasBox {
				bi = bi2
				hasBox = true
			} else if bi != bi2 {
				bi = p.size + 1
				break
			}
		}
		if !hasBox {
			return false
		}
		if last < p.Size() {
			if !p.place(r, last, v) {
				return false
			}
			continue
		}
		if bi >= p.size {
			continue
		}

		// we know that the box at bi must contain v
		// in row r. eliminate the rest.
		for bci = uint8(0); bci < p.size; bci++ {
			if r == p.boxes[bi][bci].row {
				continue
			}
			if !p.removeOption(
				p.boxes[bi][bci].row,
				p.boxes[bi][bci].col,
				b,
			) {
				return false
			}
		}
	}

	for c = 0; c < p.Size(); c++ {
		if p.placedCols[c]&b != 0 {
			continue
		}
		hasBox = false
		last = p.Size()
		for r = 0; r < p.Size(); r++ {
			if p.cannots[r][c]&b != 0 {
				// cannot play it here
				continue
			}

			if last == p.Size() {
				last = r
			} else {
				last = p.Size() + 1
			}

			bi2 = p.getBoxIndex(r, c)
			if !hasBox {
				bi = bi2
				hasBox = true
			} else if bi != bi2 {
				bi = p.size + 1
				break
			}
		}
		if !hasBox {
			return false
		}
		if last < p.Size() {
			if !p.place(last, c, v) {
				return false
			}
			continue
		}
		if bi >= p.size {
			continue
		}

		// we know that the box at bi must contain v
		// in col c. eliminate the rest.
		for bci = uint8(0); bci < p.size; bci++ {
			if c == p.boxes[bi][bci].col {
				continue
			}
			if !p.removeOption(
				p.boxes[bi][bci].row,
				p.boxes[bi][bci].col,
				b,
			) {
				return false
			}
		}
	}

	return true
}

func (p *puzzle) validate(
	v value,
) bool {
	b := v.bit()

	var bc, last boxCell

	var bci uint8

	// check each box has at least one possible cell left to place this number
	for bi := uint8(0); bi < p.size; bi++ {
		last.row = p.Size()
		for bci = uint8(0); bci < p.size; bci++ {
			bc = p.boxes[bi][bci]
			if p.grid[bc.row][bc.col] != 0 {
				if p.grid[bc.row][bc.col] == v {
					last.row = p.Size() + 1
					break
				}
			} else if p.cannots[bc.row][bc.col]&b == 0 {
				if last.row != p.Size() {
					last.row = p.Size() + 1
					break
				}
				last = bc
			}
		}
		if last.row == p.Size() {
			return false
		}
		if last.row < p.Size() && !p.place(last.row, last.col, v) {
			return false
		}
	}

	return true
}

func (p *puzzle) getBoxIndex(r, c uint8) uint8 {
	return p.box[r][c]
}

func (p *puzzle) BestRow() uint8 {
	var cur int
	br := p.Size() + 1
	b := -1

	for r := uint8(0); r < p.Size(); r++ {
		if p.remainingRows[r] == 0 {
			continue
		}
		cur = -1
		for c := uint8(0); c < p.Size(); c++ {
			if p.remaining[r][c] == 0 {
				continue
			}
			if cur == -1 {
				cur = int(p.remaining[r][c])
			} else {
				cur *= int(p.remaining[r][c])
			}
			if b > 0 && cur > b {
				break
			}
		}
		if cur < 0 {
			continue
		}
		if b < 0 || cur < b {
			b = cur
			br = r
		}
	}

	return br
}

func (p *puzzle) BestCol() uint8 {
	var cur int
	bc := p.Size() + 1
	b := -1

	for c := uint8(0); c < p.Size(); c++ {
		cur = -1
		for r := uint8(0); r < p.Size(); r++ {
			if p.remaining[r][c] == 0 {
				continue
			}
			if cur == -1 {
				cur = int(p.remaining[r][c])
			} else {
				cur *= int(p.remaining[r][c])
			}
			if b > 0 && cur > b {
				break
			}
		}
		if cur < 0 {
			continue
		}
		if b < 0 || cur < b {
			b = cur
			bc = c
		}
	}

	return bc
}

func (p *puzzle) IsSolved() bool {
	var row, col, b bits
	// check each row/col that it has all the numbers
	var r, c uint8
	for r = 0; r < p.Size(); r++ {
		if p.remainingRows[r] > 0 {
			return false
		}
		row = 0
		col = 0
		for c = 0; c < p.Size(); c++ {
			if p.grid[r][c] == 0 {
				return false
			}
			// check row
			b = p.grid[r][c].bit()
			if row&b == b {
				return false
			}
			row |= b
			// Use the (r, c) vars, but invert the order to check col
			b = p.grid[c][r].bit()
			if col&b == b {
				return false
			}
			col |= b
		}
	}

	var seen bits
	var bci uint8

	// check each box that it has all the numbers
	for bi := uint8(0); bi < p.size; bi++ {
		seen = 0
		for bci = 0; bci < p.size; bci++ {
			b = p.grid[p.boxes[bi][bci].row][p.boxes[bi][bci].col].bit()
			if seen&b == b {
				// box has already seen it
				return false
			}
			seen |= b
		}
	}

	return true
}

func (p *puzzle) String() string {
	if p.Size() == 16 {
		return p.sixteenString()
	}

	return `TODO`
}

func (p *puzzle) sixteenString() string {
	output := make([]byte, 0, p.Size()*p.Size()*2)

	for r := uint8(0); r < p.Size(); r++ {
		if r == 0 {
			output = append(output, []byte(".-------.-------.-------.-------.\n")...)
		} else if r%4 == 0 {
			output = append(output, []byte("|-------+-------+-------+-------|\n")...)

		}
		output = append(output, '|')
		for c := uint8(0); c < p.Size(); c++ {
			if p.grid[r][c] > 9 {
				output = append(output, 'A'+byte(p.grid[r][c]-10))
			} else if p.grid[r][c] > 0 {
				output = append(output, '0'+byte(p.grid[r][c]))
			} else {
				output = append(output, ' ')
			}
			if c > 0 && c%4 == 3 {
				output = append(output, '|')
			} else {
				output = append(output, ' ')
			}
		}

		output = append(output, '\n')
	}
	output = append(output, []byte("'-------'-------'-------'-------'\n")...)

	return string(output)

}
