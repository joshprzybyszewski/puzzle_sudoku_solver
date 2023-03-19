package solve

type puzzle struct {
	grid [16][16]value

	cannots [16][16]bits

	remaining     [16][16]uint8
	remainingRows [16]uint8

	size uint8
}

func NewPuzzle(
	input [][]uint8,
) puzzle {
	puzz := puzzle{
		size: uint8(len(input)),
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
			if input[r][c] == 0 {
				continue
			}
			puzz.InitialPlace(r, c, input[r][c])
		}
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

func (p *puzzle) InitialPlace(r, c, val uint8) {
	if !p.Place(r, c, value(val)) {
		panic(`dev error`)
	}
}

func (p *puzzle) placeLast(r, c uint8) (bits, bool) {
	if p.grid[r][c] != 0 {
		return 0, true
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
	placed, ok := p.place(r, c, val)
	if !ok {
		return false
	}
	if placed == 0 {
		return true
	}

	for v := value(1); v <= value(p.Size()); v++ {
		if placed&(v.bit()) == 0 {
			continue
		}

		if !p.validate(v) {
			return false
		}

		placed ^= v.bit()
		if placed == 0 {
			break
		}
	}

	return true
}

func (p *puzzle) place(r, c uint8, val value) (bits, bool) {
	if p.grid[r][c] == val {
		return 0, true
	}

	b := val.bit()
	if p.cannots[r][c]&b == b {
		return 0, false
	}

	if val > value(p.Size()) || p.grid[r][c] != 0 {
		// fmt.Printf("r: %d\n", r)
		// fmt.Printf("c: %d\n", c)
		// fmt.Printf("val: %d\n", val)
		// fmt.Printf("%s\n", p)
		panic(`dev error`)
	}

	p.grid[r][c] = val
	p.cannots[r][c] = allBits
	p.remaining[r][c] = 0
	p.remainingRows[r]--

	shouldCheckLasts := false

	// update this column (iterate through all the rows)
	for r2 := uint8(0); r2 < p.Size(); r2++ {
		if r2 == r || p.cannots[r2][c]&b == b {
			continue
		}
		if p.remaining[r2][c] == 1 {
			// This removes the last option for this cell.
			// Impossible.
			return 0, false

		}
		p.remaining[r2][c]--
		p.cannots[r2][c] |= b

		if p.remaining[r2][c] == 1 {
			shouldCheckLasts = true
		}
	}

	// update this row (iterate through all the cols)
	for c2 := uint8(0); c2 < p.Size(); c2++ {
		if c2 == c || p.cannots[r][c2]&b == b {
			continue
		}
		if p.remaining[r][c2] == 1 {
			// This removes the last option for this cell.
			// Impossible.
			return 0, false

		}
		p.remaining[r][c2]--
		p.cannots[r][c2] |= b

		if p.remaining[r][c2] == 1 {
			shouldCheckLasts = true
		}
	}

	// update this box (iterate through the 16 nearby)
	bc := p.getBoxCoords(r, c)
	for r2 := bc.startR; r2 < bc.stopR; r2++ {
		for c2 := bc.startC; c2 < bc.stopC; c2++ {
			if r2 == r && c2 == c {
				continue
			}
			if p.cannots[r2][c2]&b == b {
				continue
			}
			if p.remaining[r2][c2] == 1 {
				// This removes the last option for this cell.
				// Impossible.
				return 0, false

			}
			p.remaining[r2][c2]--
			p.cannots[r2][c2] |= b

			if p.remaining[r2][c2] == 1 {
				shouldCheckLasts = true
			}
		}
	}

	out := val.bit()
	if shouldCheckLasts {
		placed, success := p.checkAllForLast()
		if !success {
			return 0, false
		}
		out |= placed
	}

	return out, true
}

func (p *puzzle) checkAllForLast() (bits, bool) {
	var c uint8
	var placed, tmp bits
	var success bool

	for r := uint8(0); r < p.Size(); r++ {
		for c = uint8(0); c < p.Size(); c++ {
			if p.remaining[r][c] != 1 {
				continue
			}

			tmp, success = p.placeLast(r, c)
			if !success {
				return 0, false
			}
			placed |= tmp
		}
	}

	return placed, true
}

func (p *puzzle) validate(
	v value,
) bool {
	b := v.bit()

	var canRow, canCol bool
	var r, c uint8

	// Check that each row and each col has at least one possible cell left for placing this value
	for r = 0; r < p.Size(); r++ {
		canRow, canCol = false, false
		for c = 0; c < p.Size(); c++ {
			// check row
			if p.grid[r][c] != 0 {
				if p.grid[r][c] == v {
					canRow = true
					if canCol {
						break
					}
				}
			} else if p.cannots[r][c]&b == 0 {
				canRow = true
				if canCol {
					break
				}
			}
			// Use the (r, c) vars, but invert the order to check col
			if p.grid[c][r] != 0 {
				if p.grid[c][r] == v {
					canCol = true
					if canRow {
						break
					}
				}
			} else if p.cannots[c][r]&b == 0 {
				canCol = true
				if canRow {
					break
				}
			}
		}
		if !canRow || !canCol {
			return false
		}
	}

	var canBox bool
	// check each box has at least one possible cell left to place this number
	for bc := p.getBoxCoords(0, 0); ; {
		canBox = false
		for r = bc.startR; r < bc.stopR; r++ {
			for c = bc.startC; c < bc.stopC; c++ {
				if p.grid[r][c] != 0 {
					if p.grid[r][c] == v {
						canBox = true
						break
					}
				} else if p.cannots[r][c]&b == 0 {
					canBox = true
					break
				}
			}
			if canBox {
				break
			}
		}
		if !canBox {
			return false
		}
		if bc.stopC < p.Size() {
			bc = p.getBoxCoords(bc.startR, bc.stopC)
		} else if bc.stopR < p.Size() {
			bc = p.getBoxCoords(bc.stopR, 0)
		} else {
			break
		}
	}

	return true
}

func (p *puzzle) getBoxCoords(r, c uint8) boxCoords {
	switch p.size {
	case 9:
		return nineBoxCoords[r][c]
	case 12:
		return twelveBoxCoords[r][c]
	case 16:
		return sixteenBoxCoords[r][c]
	}
	return invalidBoxCoords
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
			br = uint8(r)
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
	var row, col, box, b bits
	// check each row/col that it has all the numbers
	var r, c uint8
	for r = 0; r < p.Size(); r++ {
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

	// check each box that it has all the numbers
	for bc := p.getBoxCoords(0, 0); ; {
		box = 0
		for r = bc.startR; r < bc.stopR; r++ {
			for c = bc.startC; c < bc.stopC; c++ {
				b = p.grid[r][c].bit()
				if box&b == b {
					// box has already seen it
					return false
				}
				box |= b
			}
		}
		if bc.stopC < p.Size() {
			bc = p.getBoxCoords(bc.startR, bc.stopC)
		} else if bc.stopR < p.Size() {
			bc = p.getBoxCoords(bc.stopR, 0)
		} else {
			break
		}
	}

	return true
}

func (p puzzle) String() string {
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
