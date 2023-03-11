package solve

import (
	"errors"
	"time"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
)

func Sixteen(
	puzzle model.Sixteen,
	timeout time.Duration,
) (model.Sixteen, error) {
	s, ok := solveSixteen(puzzle)
	if !ok {
		return model.Sixteen{}, errors.New(`did not solve`)
	}
	return s, nil
}

func solveSixteen(
	s model.Sixteen,
) (model.Sixteen, bool) {
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				continue
			}

			for i := 1; i <= len(s); i++ {
				s[r][c] = uint8(i)
				if !isValidSixteen(s) {
					continue
				}
				solved, ok := solveSixteen(s)
				if ok {
					return solved, true
				}
			}

			return model.Sixteen{}, false
		}
	}

	return s, isValidSixteen(s)
}

func isValidSixteen(
	p model.Sixteen,
) bool {
	var seen, b uint16
	// check each row that it has all the numbers
	for r := 0; r < len(p); r++ {
		seen = 0
		for c := range p[r] {
			if p[r][c] == 0 {
				continue
			}
			b = 1 << (p[r][c] - 1)
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// check each col that it has all the numbers
	for c := 0; c < len(p[0]); c++ {
		seen = 0
		for r := 0; r < len(p); r++ {
			if p[r][c] == 0 {
				continue
			}
			b = 1 << (p[r][c] - 1)
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// check each box that it has all the numbers
	for box := 0; box < len(p[0]); box++ {
		seen = 0
		for r := 4 * (box / 4); r < 4*(box/4)+4; r++ {
			for c := 4 * (box % 4); c < 4*(box%4)+4; c++ {
				if p[r][c] == 0 {
					continue
				}
				b = 1 << (p[r][c] - 1)
				if seen&b == b {
					return false
				}
				seen |= b
			}
		}
	}

	return true
}
