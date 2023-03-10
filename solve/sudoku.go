package solve

import (
	"errors"
	"time"

	"github.com/joshprzybyszewski/puzzle_sudoku_solver/model"
)

func SolveClassic(
	puzzle model.Classic,
	timeout time.Duration,
) (model.Classic, error) {
	s, ok := solve(puzzle)
	if !ok {
		return model.Classic{}, errors.New(`did not solve`)
	}
	return s, nil
}

func solve(
	s model.Classic,
) (model.Classic, bool) {
	for r := range s {
		for c := range s[r] {
			if s[r][c] != 0 {
				continue
			}

			for i := 1; i <= len(s); i++ {
				s[r][c] = uint8(i)
				if !isValid(s) {
					continue
				}
				solved, ok := solve(s)
				if ok {
					return solved, true
				}
			}

			return model.Classic{}, false
		}
	}

	return s, isValid(s)
}

func isValid(
	p model.Classic,
) bool {
	var seen, b uint16
	// check each row that it has all the numbers
	for r := 0; r < len(p); r++ {
		seen = 0
		for c := range p[r] {
			if p[r][c] == 0 {
				continue
			}
			b = 1 << p[r][c]
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
			b = 1 << p[r][c]
			if seen&b == b {
				return false
			}
			seen |= b
		}
	}

	// TODO
	// check each box that it has all the numbers
	for box := 0; box < len(p[0]); box++ {
		seen = 0
		for r := 3 * (box / 3); r < 3*(box/3)+3; r++ {
			for c := 3 * (box % 3); c < 3*(box%3)+3; c++ {
				if p[r][c] == 0 {
					continue
				}
				b = 1 << p[r][c]
				if seen&b == b {
					return false
				}
				seen |= b
			}
		}
	}

	return true
}
