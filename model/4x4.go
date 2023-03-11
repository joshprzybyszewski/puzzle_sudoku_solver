package model

type Sixteen struct {
	grid [16][16]uint8
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
	p.place(r, c, val)
}

func (p *Sixteen) place(r, c, val uint8) {
	p.grid[r][c] = val
}

func (p *Sixteen) Place(r, c, val uint8) (Sixteen, bool) {
	// check that this row doesn't have val
	for c := range p.grid[r] {
		if p.grid[r][c] == val {
			return Sixteen{}, false

		}
	}

	// check that this column doesn't have val
	for r := range p.grid {
		if p.grid[r][c] == val {
			return Sixteen{}, false
		}
	}

	// check each box that it has all the numbers
	startR := 4 * (r / 4)
	stopR := startR + 4
	startC := 4 * (c / 4)
	stopC := startC + 4
	for r := startR; r < stopR; r++ {
		for c := startC; c < stopC; c++ {
			if p.grid[r][c] == val {
				return Sixteen{}, false
			}
		}
	}

	cpy := *p
	cpy.place(r, c, val)
	return cpy, true
}

func (p *Sixteen) IsSolved() bool {
	var seen, b uint16
	// check each row that it has all the numbers
	for r := 0; r < len(p.grid); r++ {
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
	for c := 0; c < len(p.grid[0]); c++ {
		seen = 0
		for r := 0; r < len(p.grid); r++ {
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
		for c := range p.grid[r] {
			if p.grid[r][c] > 9 {
				output = append(output, 'A'+(p.grid[r][c]-9))
			} else {
				output = append(output, '0'+p.grid[r][c])
			}
			output = append(output, ',')
		}
		output = append(output, '\n')
	}

	// omit the last comma
	return string(output[:len(output)-1])

}
