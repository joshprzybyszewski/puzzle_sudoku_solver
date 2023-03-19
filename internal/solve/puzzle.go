package solve

import "github.com/joshprzybyszewski/puzzle_sudoku_solver/internal/model"

type puzzle struct {
	grid [16][16]value

	cannots [16][16]bits

	remaining     [16][16]uint8
	remainingRows [16]uint8

	size uint8
}

func NewPuzzleFromSixteen(
	input model.Sixteen,
) puzzle {
	puzz := puzzle{
		size: 16,
	}

	for r := range puzz.remaining {
		for c := range puzz.remaining[r] {
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
	return uint8(len(p.grid))
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

	return p.Place(r, c, val)
}

func (p *puzzle) Place(r, c uint8, val value) bool {
	if p.grid[r][c] == val {
		return true
	}

	b := val.bit()
	if p.cannots[r][c]&b == b {
		return false
	}

	if val > value(p.Size()) || p.grid[r][c] != 0 {
		// fmt.Printf("r: %d\n", r)
		// fmt.Printf("c: %d\n", c)
		// fmt.Printf("val: %d\n", val)
		// fmt.Printf("%s\n", p)
		panic(`dev error`)
	}

	p.grid[r][c] = val
	// p.cannots[r][c] = 0xFFFF
	p.remaining[r][c] = 0
	p.remainingRows[r]--

	shouldCheckLasts := false

	// update this column (iterate through all the rows)
	for r2 := uint8(0); r2 < p.Size(); r2++ {
		if r2 == r || p.remaining[r2][c] == 0 || p.cannots[r2][c]&b == b {
			continue
		}
		if p.remaining[r2][c] == 1 {
			return false
		}
		p.remaining[r2][c]--
		p.cannots[r2][c] |= b

		if p.remaining[r2][c] == 1 {
			shouldCheckLasts = true
		}
	}

	// update this row (iterate through all the cols)
	for c2 := uint8(0); c2 < p.Size(); c2++ {
		if c2 == c || p.remaining[r][c2] == 0 || p.cannots[r][c2]&b == b {
			continue
		}
		if p.remaining[r][c2] == 1 {
			return false
		}
		p.remaining[r][c2]--
		p.cannots[r][c2] |= b

		if p.remaining[r][c2] == 1 {
			shouldCheckLasts = true
		}
	}

	// update this box (iterate through the 16 nearby)
	startR := 4 * (r / 4)
	stopR := startR + 4
	startC := 4 * (c / 4)
	stopC := startC + 4
	for r2 := startR; r2 < stopR; r2++ {
		for c2 := startC; c2 < stopC; c2++ {
			if r2 == r && c2 == c {
				continue
			}
			if p.remaining[r2][c2] == 0 || p.cannots[r2][c2]&b == b {
				continue
			}
			if p.remaining[r2][c2] == 1 {
				return false
			}
			p.remaining[r2][c2]--
			p.cannots[r2][c2] |= b

			if p.remaining[r2][c2] == 1 {
				shouldCheckLasts = true
			}
		}
	}

	if !shouldCheckLasts {
		return true
	}

	for r := uint8(0); r < p.Size(); r++ {
		for c := uint8(0); c < p.Size(); c++ {
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

func (p *puzzle) BestRow() uint8 {
	var cur int
	br := p.Size() + 1
	b := -1

	for r := range p.remaining {
		if p.remainingRows[r] == 0 {
			continue
		}
		cur = -1
		for c := range p.remaining[r] {
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
	var seen, b uint16
	// check each row that it has all the numbers
	for r := uint8(0); r < p.Size(); r++ {
		seen = 0
		for c := range p.grid[r] {
			if p.grid[r][c] == 0 {
				return false
			}
			b = 1 << (p.grid[r][c] - 1)
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// check each col that it has all the numbers
	for c := uint8(0); c < p.Size(); c++ {
		seen = 0
		for r := uint8(0); r < p.Size(); r++ {
			b = 1 << (p.grid[r][c] - 1)
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// check each box that it has all the numbers
	for box := 0; box < len(p.grid[0]); box++ {
		seen = 0
		for r := 4 * (box / 4); r < 4*(box/4)+4; r++ {
			for c := 4 * (box % 4); c < 4*(box%4)+4; c++ {
				b = 1 << (p.grid[r][c] - 1)
				if seen&b == b {
					return false
				}
				seen |= b
			}
		}
	}

	return true
}

func (p puzzle) String() string {
	output := make([]byte, 0, len(p.grid)*len(p.grid[0])*2)

	for r := range p.grid {
		if r == 0 {
			output = append(output, []byte(".-------.-------.-------.-------.\n")...)
		} else if r%4 == 0 {
			output = append(output, []byte("|-------+-------+-------+-------|\n")...)

		}
		output = append(output, '|')
		for c := range p.grid[r] {
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

		// for c := 0; c < 4; c++ {
		// 	output = append(output, []byte(fmt.Sprintf(" %016b", p.cannots[r][c]))...)
		// }
		output = append(output, '\n')
	}
	output = append(output, []byte("'-------'-------'-------'-------'\n")...)

	return string(output)

}
