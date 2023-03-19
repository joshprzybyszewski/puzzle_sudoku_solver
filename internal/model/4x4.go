package model

import (
	pmodel "github.com/joshprzybyszewski/puzzler/model"
)

type Sixteen struct {
	grid [16][16]uint8

	cannots [16][16]uint16

	remaining     [16][16]uint8
	remainingRows [16]uint8
}

func NewSixteen(
	task pmodel.Task,
) Sixteen {
	puzz := Sixteen{}

	for r := range puzz.remaining {
		for c := range puzz.remaining[r] {
			puzz.remaining[r][c] = puzz.Size()
		}
		puzz.remainingRows[r] = puzz.Size()
	}

	var r, c uint8

	var b byte
	for i := 0; i < len(task); i++ {
		b = task[i]
		if b == '_' {
			continue
		}

		if b >= '0' && b <= '9' {
			if b == '1' && i+1 < len(task) && task[i+1] >= '0' && task[i+1] <= '6' {
				puzz.InitialPlace(r, c, 10+uint8(task[i+1]-'0'))
				i++
			} else {
				puzz.InitialPlace(r, c, uint8(b-'0'))
			}
		} else {
			c += uint8(b - 'a')
		}

		c++

		if c >= puzz.Size() {
			r += (c / puzz.Size())
			c %= puzz.Size()
		}
	}

	return puzz
}

func (p *Sixteen) Size() uint8 {
	return uint8(len(p.grid))
}

func (p *Sixteen) Grid() [16][16]uint8 {
	return p.grid
}

func (p *Sixteen) IsSet(r, c uint8) bool {
	return p.grid[r][c] != 0
}

func (p *Sixteen) CanPlace(r, c, val uint8) bool {
	b := uint16(1) << (val - 1)
	return p.cannots[r][c]&b == 0
}

func (p *Sixteen) Val(r, c uint8) uint8 {
	return p.grid[r][c]
}

func (p *Sixteen) InitialPlace(r, c, val uint8) {
	if !p.Place(r, c, val) {
		panic(`dev error`)
	}
}

func (p *Sixteen) placeLast(r, c uint8) bool {
	if p.grid[r][c] != 0 {
		return true
	}
	if p.remaining[r][c] != 1 {
		panic(`dev error`)
	}
	val := uint8(1)
	b := uint16(1)
	for p.cannots[r][c]&b == b {
		val++
		b <<= 1
	}

	return p.Place(r, c, val)
}

func (p *Sixteen) Place(r, c, val uint8) bool {
	if p.grid[r][c] == val {
		return true
	}

	b := uint16(1) << (val - 1)
	if p.cannots[r][c]&b == b {
		return false
	}

	if val > p.Size() || p.grid[r][c] != 0 {
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

func (p *Sixteen) BestRow() uint8 {
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

func (p *Sixteen) BestCol() uint8 {
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

		// for c := 0; c < 4; c++ {
		// 	output = append(output, []byte(fmt.Sprintf(" %016b", p.cannots[r][c]))...)
		// }
		output = append(output, '\n')
	}
	output = append(output, []byte("'-------'-------'-------'-------'\n")...)

	return string(output)

}
