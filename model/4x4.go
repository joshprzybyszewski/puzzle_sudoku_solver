package model

import "fmt"

type Sixteen struct {
	grid [16][16]uint8

	cannots [16][16]uint16

	remaining     [16][16]uint8
	remainingRows [16]uint8
}

func NewSixteen() Sixteen {
	s := Sixteen{}

	for r := range s.remaining {
		for c := range s.remaining[r] {
			s.remaining[r][c] = s.Size()
		}
		s.remainingRows[r] = s.Size()
	}

	return s
}

func (p *Sixteen) Size() uint8 {
	return uint8(len(p.grid))
}

func (p *Sixteen) IsSet(r, c uint8) bool {
	return p.grid[r][c] != 0
}

func (p *Sixteen) Val(r, c uint8) uint8 {
	return p.grid[r][c]
}

func (p *Sixteen) InitialPlace(r, c, val uint8) {
	if !p.place(r, c, val) {
		panic(`dev error`)
	}
}

func (p *Sixteen) placeLast(r, c uint8) bool {
	if p.grid[r][c] != 0 || p.remaining[r][c] != 1 {
		panic(`dev error`)
	}
	val := uint8(1)
	b := uint16(1)
	for p.cannots[r][c]&b == b {
		val++
		b <<= 1
	}

	return p.place(r, c, val)
}

func (p *Sixteen) place(r, c, val uint8) bool {

	b := uint16(1) << (val - 1)
	if p.cannots[r][c]&b == b {
		return false
	}

	if val > p.Size() || p.grid[r][c] != 0 {
		panic(`dev error`)
	}

	p.grid[r][c] = val
	p.cannots[r][c] = 0xFFFF
	p.remaining[r][c] = 0
	p.remainingRows[r]--

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
			if !p.placeLast(r2, c) {
				return false
			}
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
			if !p.placeLast(r, c2) {
				return false
			}
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
				if !p.placeLast(r2, c2) {
					return false
				}
			}
		}
	}

	return true
}

func (p *Sixteen) Place(r, c, val uint8) (Sixteen, bool) {
	cpy := *p
	return cpy, cpy.place(r, c, val)
}

func (p *Sixteen) Best() (uint8, uint8) {
	var r uint8
	b := p.Size() + 1
	for other := uint8(0); other < p.Size(); other++ {
		if p.remainingRows[other] == 1 {
			r = other
			break
		}
		if p.remainingRows[other] > 0 && p.remainingRows[other] < b {
			b = p.remainingRows[other]
			r = other
		}
	}

	if b > p.Size() {
		// did not find!
		return p.Size() + 1, p.Size() + 1
	}

	b = p.Size() + 1
	var c uint8

	for j := range p.remaining[r] {
		if p.remaining[r][j] == 1 {
			return r, uint8(j)
		}
		if p.remaining[r][j] > 0 && p.remaining[r][j] < b {
			b = p.remaining[r][j]
			c = uint8(j)
		}
	}

	if b > p.Size() {
		// did not find!
		return p.Size() + 1, p.Size() + 1
	}

	return r, c
}

func (p *Sixteen) IsSolved() bool {
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

	// check each box that it has all the numbers
	for box := 0; box < len(p.grid[0]); box++ {
		seen = 0
		for r := 4 * (box / 4); r < 4*(box/4)+4; r++ {
			for c := 4 * (box % 4); c < 4*(box%4)+4; c++ {
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
	}

	return true
}

func (p Sixteen) String() string {
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
				output = append(output, 'A'+(p.grid[r][c]-10))
			} else if p.grid[r][c] > 0 {
				output = append(output, '0'+p.grid[r][c])
			} else {
				output = append(output, ' ')
			}
			if c > 0 && c%4 == 3 {
				output = append(output, '|')
			} else {
				output = append(output, ' ')
			}
		}

		for c := 0; c < 4; c++ {
			output = append(output, []byte(fmt.Sprintf(" %016b", p.cannots[r][c]))...)
		}
		output = append(output, '\n')
	}
	output = append(output, []byte("'-------'-------'-------'-------'\n")...)

	return string(output)

}
